package uuid

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// New 生成一个随机的唯一 ID
func New() string {
	return uuid.NewV4().String()
}

func UUID16() string {
	uuidStr := strings.ReplaceAll(New(), "-", "")
	return uuidStr[0:16]
}
