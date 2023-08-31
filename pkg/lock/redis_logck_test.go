package lock

// import (
// 	"context"
// 	"sync"
// 	"testing"
// 	"time"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSimpleLock_LockUser(t *testing.T) {
// 	ctx := context.Background()

// 	locker := getLocker()
// 	locker.LockUser(1234567890)
// 	value := locker.redisClient.Get(ctx, "user_distributed_lock_key_1234567890").Val()
// 	assert.Equal(t, value, locker.uuid, "")

// 	re := locker.UnlockUser(1234567890)
// 	assert.True(t, re)
// 	value = locker.redisClient.Get(ctx, "user_distributed_lock_key_1234567890").Val()
// 	assert.Equal(t, value, "", "")
// }

// func TestSimpleLock_LockUser2(t *testing.T) {
// 	ctx := context.Background()

// 	var locker *RedisLocker
// 	func() {
// 		locker = getLocker()
// 		locker.LockUser(1234567890)
// 		defer locker.UnlockUser(1234567890)
// 		value := locker.redisClient.Get(ctx, "user_distributed_lock_key_1234567890").Val()
// 		assert.Equal(t, value, locker.uuid, "")
// 	}()
// 	value := locker.redisClient.Get(ctx, "user_distributed_lock_key_1234567890").Val()
// 	assert.Equal(t, value, "", "")
// }

// func TestSimpleLock_Lock(t *testing.T) {
// 	ctx := context.Background()

// 	distLock := getLocker()
// 	distLock.Lock("test")
// 	value := distLock.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 	assert.Equal(t, value, distLock.uuid, "")

// 	re := distLock.Unlock("test")
// 	assert.True(t, re)
// 	value = distLock.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 	assert.Equal(t, value, "", "")
// }

// func TestNewRedisLock(t *testing.T) {
// 	var count int
// 	var wg sync.WaitGroup
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			distLock := getLocker()
// 			distLock.Lock("test")
// 			for i := 0; i < 1000; i++ {
// 				count = count + 100
// 				time.Sleep(1000)
// 			}
// 			distLock.Unlock("test")
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// 	assert.Equal(t, count, 1000000, "")
// }

// func TestNewUserLock(t *testing.T) {
// 	var count1 int
// 	var count2 int
// 	var wg sync.WaitGroup

// 	wg.Add(20)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			distLock := getLocker()
// 			distLock.LockUser(1234567890)
// 			for i := 0; i < 1000; i++ {
// 				count1 = count1 + 100
// 				time.Sleep(1000)
// 			}
// 			distLock.UnlockUser(1234567890)
// 			wg.Done()
// 		}()

// 		go func() {
// 			distLock := getLocker()
// 			distLock.LockUser(1111111110)
// 			for i := 0; i < 1000; i++ {
// 				count2 = count2 + 99
// 				time.Sleep(1000)
// 			}
// 			distLock.UnlockUser(1111111110)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// 	assert.Equal(t, count1, 1000000, "")
// 	assert.Equal(t, count2, 990000, "")
// 	//t.Log(count1)
// 	//t.Log(count2)
// }

// func TestABA(t *testing.T) {
// 	ctx := context.Background()

// 	distLock1 := getLocker(WithExpiry(time.Second*5), WithAutoRenewal(true))
// 	assert.True(t, distLock1.Lock("test"))
// 	value := distLock1.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 	assert.Equal(t, value, distLock1.uuid, "")

// 	time.Sleep(time.Second * 5)

// 	// distLock2 并发请求加锁
// 	distLock2 := getLocker()

// 	// distLock1 因为有锁补偿延期机制，不会导致锁超时，所以这里 distLock2 加锁不应该成功
// 	assert.False(t, distLock2.Lock("test"))
// 	value = distLock2.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 	assert.Equal(t, value, distLock1.uuid, "")

// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		time.Sleep(time.Second)

// 		// distLock2 等待 distLock1 释放锁后成功加锁
// 		assert.True(t, distLock2.Lock("test"))
// 		// NOTE: 这里要用 := 否则会被抓住 data race
// 		value := distLock2.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 		assert.Equal(t, value, distLock2.uuid, "")

// 		assert.True(t, distLock2.Unlock("test"))
// 		value = distLock2.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 		assert.Equal(t, value, "", "")

// 		wg.Done()
// 	}()

// 	// 释放 distLock1 锁以让 distLock2 成功加锁
// 	assert.True(t, distLock1.Unlock("test"))
// 	value = distLock1.redisClient.Get(ctx, "simple_distributed_lock_key_test").Val()
// 	assert.Equal(t, value, "", "")

// 	wg.Wait()
// }

// func getLocker(opts ...RedisLockerOption) *RedisLocker {
// 	return NewRedisLocker(getRedisClient(), opts...).(*RedisLocker)
// }

// func getRedisClient() *redis.Client {
// 	return redis.NewClient(&redis.Options{
// 		Addr: "127.0.0.1:6379",
// 	})
// }
