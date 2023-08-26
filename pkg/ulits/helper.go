package ulits

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net/smtp"
	"sort"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
)

const DATE_ONLY = "2006-01-02"

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

// TimeSort 时间排序
func TimeSort(times []string, format string) []string {

	sort.Slice(times, func(i, j int) bool {
		t1, err := time.Parse(format, times[i])
		if err != nil {
			log.Printf("time.Parse(%s,%s) faild err: %v\n", format, times[i], err)
			return false
		}

		t2, err := time.Parse(format, times[j])
		if err != nil {
			log.Printf("time.Parse(%s,%s) faild err: %v\n", format, times[j], err)
			return false
		}

		return t1.Before(t2)
	})

	return times
}
