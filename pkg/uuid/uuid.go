package uuid

import (
	"strings"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	letters    = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	latestTime int64
	count      int64
	idMutex    sync.Mutex
)

func ID() uint64 {
	idMutex.Lock()
	defer idMutex.Unlock()
	nowTime := time.Now().UnixNano() / 1e10
	if latestTime == nowTime {
		count++
	} else {
		latestTime = nowTime
		count = 0
	}
	res := nowTime
	res <<= 15
	res += count
	return uint64(res)
}

// New 生成一个随机的唯一 ID
func New() string {
	return uuid.NewV4().String()
}

func UUID16() string {
	uuidStr := strings.ReplaceAll(New(), "-", "")
	return uuidStr[0:16]
}
