package lock

// Locker 通用的锁接口
type Locker interface {
	Lock(key string) bool       // Lock 根据指定的 key 锁定资源
	Unlock(key string) bool     // Unlock 根据指定的 key 解锁资源
	LockUser(uid uint64) bool   // LockUser 锁定用户，相当于对用户加了全局锁
	UnlockUser(uid uint64) bool // UnlockUser 解锁用户
}
