package ulits

import (
	"crypto/tls"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	UUID     string `json:"uuid"`
	UserName string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
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

// AnalyseToken 解析 token
func AnalyseToken(tokenString string) (*UserInfo, error) {
	UserInfo := &UserInfo{}
	claims, err := jwt.ParseWithClaims(tokenString, UserInfo, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if err != nil || !claims.Valid {
		return nil, err
	}
	return UserInfo, nil
}

// SendMail 发送邮件
func SendMail(toEmail string, content []byte) error {
	e := email.NewEmail()
	e.From = "Get <test@163.com>"
	e.To = []string{toEmail}
	e.Subject = "gin-demo"
	e.HTML = content
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "test@163.com", "password", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}

// GetCode 生成验证码
func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
