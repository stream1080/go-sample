package controller

import (
	"log"
	"time"

	"go-sample/models"
	"go-sample/pkg/ulits"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserApi struct {
}

// SendCode
// @Tags 公共方法
// @Summary 发送邮件验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send/code [post]
func (u *UserApi) SendCode(c *gin.Context) {

	email := c.PostForm("email")
	if email == "" {
		ResponseError(c, InvalidArgs)
		return
	}
	code := ulits.GetCode()
	_, err := models.Redis.Set(email, code, 5*60*time.Minute).Result()
	if err != nil {
		log.Printf("Set Code Error:%v \n", err)
		ResponseError(c, ServerError)
		return
	}
	content := []byte("您的验证码为：" + code + ", 5分钟内有效，请及时操作。")
	ulits.SendMail(email, content)

	ResponseSuccess(c, nil)
}

// Register
// @Tags 用户管理
// @Summary 用户注册
// @Param code formData string true "code"
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param mobile formData string false "mobile"
// @Success 200 {string} json "{"code":"200","msg":"success","data":""}"
// @Router /register [post]
func (u *UserApi) Register(c *gin.Context) {
	var (
		username = c.PostForm("username")
		password = c.PostForm("password")
		mobile   = c.PostForm("mobile")
		code     = c.PostForm("code")
	)
	if code == "" || username == "" || password == "" {
		ResponseError(c, InvalidArgs)
		return
	}
	// 验证验证码是否正确
	verificationCode, err := models.Redis.Get(mobile).Result()
	if err != nil {
		log.Printf("Get Code Error:%v \n", err)
		ResponseError(c, CodeExpire)
		return
	}
	if verificationCode != code {
		ResponseError(c, CodeError)
		return
	}
	// 判断邮箱是否已存在
	var cnt int64
	err = models.DB.Where("mobile = ?", mobile).Model(new(models.User)).Count(&cnt).Error
	if err != nil {
		ResponseError(c, ServerError)
		return
	}
	if cnt > 0 {
		ResponseError(c, UserExist)
		return
	}

	// 数据的插入
	uuid := ulits.GetUUID()
	user := &models.User{
		UUID:     uuid,
		UserName: username,
		Password: ulits.GetMd5(password),
		Mobile:   mobile,
	}
	err = models.DB.Create(user).Error
	if err != nil {
		ResponseError(c, ServerError)
		return
	}

	// 生成 token
	token, err := ulits.GetToken(user.UUID, user.UserName, user.Role)
	if err != nil {
		ResponseError(c, ServerError)
		return
	}

	data := map[string]string{
		"token": token,
	}

	ResponseSuccess(c, data)
}

// Login
// @Tags 用户管理
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":"","msg":"success"}"
// @Router /login [post]
func (u *UserApi) Login(c *gin.Context) {
	user := new(models.User)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		ResponseError(c, InvalidArgs)
		return
	}
	// password = ulits.Md5(password)

	if err := models.DB.Where("username = ? and password = ?", username, password).First(&user).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			ResponseError(c, InvalidArgs)
			return
		}

		log.Printf("Models err: %s", err)
		ResponseError(c, ServerError)

		return
	}

	token, err := ulits.GetToken(user.UUID, user.UserName, user.Role)
	if err != nil {
		ResponseError(c, Unauthorized)
		log.Printf("GetToken err: %s", err)
		return
	}

	data := map[string]string{
		"token": token,
	}

	ResponseSuccess(c, data)
}

// GetUserInfo
// @Tags 用户管理
// @Summary 用户详情
// @Param authorization header string true "authorization"
// @Param id query string false "id"
// @Success 200 {string} json "{"code":"200","data":"","msg":"success"}"
// @Router /user/info [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		ResponseError(c, InvalidArgs)
		return
	}
	data := new(models.User)
	err := models.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		ResponseError(c, InvalidArgs)
		return
	}

	ResponseSuccess(c, data)
}
