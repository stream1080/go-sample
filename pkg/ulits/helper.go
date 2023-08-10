package ulits

import (
	"crypto/tls"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

// GetUUID 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
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
