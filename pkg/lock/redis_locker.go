package lock

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// RedisLocker 基于 redis 实现的简单分布式锁
type RedisLocker struct {
	uuid         string                //uuid 加锁的客户端 Id
	expiry       time.Duration         // expiry 锁的过期时间
	spinTimes    int                   // spinTimes 获取锁的最大尝试次数
	spinInterval time.Duration         // spinInterval 尝试获取锁的时间间隔
	cancelFunc   context.CancelFunc    // cancelFunc 用于取消补偿延期的协程
	autoRenewal  bool                  // autoRenewal 是否自动延期持有锁
	redisClient  redis.UniversalClient // redis 客户端
}

const defaultExpiry = 60 * 60 * time.Second
const defaultSpinTimes = 5000
const defaultSpinInterval = 1 * time.Millisecond

// RedisLockerOption options
type RedisLockerOption func(*RedisLocker)

// NewRedisLocker 新建一个基于 redis 的分布式锁
func NewRedisLocker(client redis.UniversalClient, opts ...RedisLockerOption) Locker {
	locker := &RedisLocker{
		uuid:         uuid.NewString(),
		redisClient:  client,
		expiry:       defaultExpiry,
		spinTimes:    defaultSpinTimes,
		spinInterval: defaultSpinInterval,
		autoRenewal:  false, // 自动续期
	}

	for _, opt := range opts {
		opt(locker)
	}

	return locker
}

// WithExpiry 设置锁过期时间
func WithExpiry(expiry time.Duration) RedisLockerOption {
	return func(locker *RedisLocker) {
		// 不允许设置永不过期
		if expiry > 0 {
			locker.expiry = expiry
		}
	}
}

// WithSpinTimes 设置尝试获取锁的次数
func WithSpinTimes(spinTimes int) RedisLockerOption {
	return func(locker *RedisLocker) {
		// 尝试次数必须大于 0
		if spinTimes > 0 {
			locker.spinTimes = spinTimes
		}
	}
}

// WithSpinInterval 设置尝试获取锁的间隔
func WithSpinInterval(spinInterval time.Duration) RedisLockerOption {
	return func(locker *RedisLocker) {
		// 间隔时间必须大于 0
		if spinInterval > 0 {
			locker.spinInterval = spinInterval
		}
	}
}

// WithAutoRenewal 设置是否自动延期持有锁
func WithAutoRenewal(autoRenewal bool) RedisLockerOption {
	return func(locker *RedisLocker) {
		locker.autoRenewal = autoRenewal
	}
}

// LockUser 业务层的用户锁，锁定用户，业务层的排他锁
func (locker *RedisLocker) LockUser(uid uint64) bool {
	return locker.lock(getUserKey(uid))
}

// Lock 分布式锁的基础接口，对指定 key 加锁
func (locker *RedisLocker) Lock(key string) bool {
	return locker.lock(getLockKey(key))
}

func (locker *RedisLocker) lock(key string) bool {
	i := 0
	for i < locker.spinTimes {
		result := locker.redisClient.SetNX(context.TODO(), key, locker.uuid, locker.expiry)
		if result.Val() && result.Err() == nil {
			if locker.autoRenewal {
				ctx, cancelFunc := context.WithCancel(context.TODO())
				locker.cancelFunc = cancelFunc
				locker.renew(ctx, key, locker.expiry)
			}

			return true
		}

		time.Sleep(locker.spinInterval)
		i++
	}

	return false
}

// UnlockUser 释放业务层的用户锁
func (locker *RedisLocker) UnlockUser(uid uint64) bool {
	return locker.unlock([]string{getUserKey(uid)})
}

// Unlock 分布式锁的基础接口，对指定的 key 解锁
func (locker *RedisLocker) Unlock(key string) bool {
	return locker.unlock([]string{getLockKey(key)})
}

func (locker *RedisLocker) unlock(keys []string) bool {
	result := locker.redisClient.Eval(context.TODO(), script, keys, locker.uuid)

	//result.Val() == (int64(1)) 自己删除key
	//result.Val() == (int64(0)) 这个key已经过期，找不到了
	isSucc := result.Err() == nil
	if isSucc && locker.cancelFunc != nil {
		locker.cancelFunc()
	}

	return isSucc
}

func (locker *RedisLocker) renew(ctx context.Context, key string, expiry time.Duration) {
	go func() {
		tick := time.NewTicker(expiry / 3)
		for {
			select {
			case <-ctx.Done():
				tick.Stop()
				return
			case <-tick.C:
				locker.redisClient.Eval(
					context.TODO(),
					renewalScript,
					[]string{key},
					locker.uuid,
					int64(expiry/time.Second),
				)
			}
		}
	}()
}

func getUserKey(uid uint64) string {
	return "user_distributed_lock_key_" + strconv.FormatInt(int64(uid), 10)
}

func getLockKey(key string) string {
	return "simple_distributed_lock_key_" + key
}

const script = `
	if redis.call('get',KEYS[1]) == ARGV[1] then 
		return redis.call('del',KEYS[1]) 
	else 
		return 0 
	end
	`
const renewalScript = `
	if redis.call('get',KEYS[1]) == ARGV[1] then 
		return redis.call('expire',KEYS[1],ARGV[2]) 
	else 
		return 0 
	end
	`
