package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5WithSalt 加密时简单加盐
func Md5Salt(src string, salt string) string {
	return Md5(src + "#" + salt)
}

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}
