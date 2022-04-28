package ulits

import (
	"crypto/md5"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

// GetMd5 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GetUUID 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}

var tokenKey = []byte("test-key")

// GetToken 生成 token
func GetToken(uuid, username string, role int) (string, error) {
	UserInfo := UserInfo{
		UUID:           uuid,
		UserName:       username,
		Role:           role,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserInfo)
	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
