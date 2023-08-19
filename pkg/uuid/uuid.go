package uuid

import (
	uuid "github.com/satori/go.uuid"
)

// New 生成一个随机的唯一 ID
func New() string {
	return uuid.NewV4().String()
}
