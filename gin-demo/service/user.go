package service

import (
	"gin-demo/api"
	"gin-demo/models"
	"gin-demo/ulits"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendCode
// @Tags 公共方法
// @Summary 发送邮件验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send/code [post]
func SendCode(c *gin.Context) {
	res := api.NewResult(c)
	email := c.PostForm("email")
	if email == "" {
		res.Error(api.InvalidArgs)
		return
	}
	code := ulits.GetCode()
	_, err := models.Redis.Set(email, code, 5*60*time.Minute).Result()
	if err != nil {
		log.Printf("Set Code Error:%v \n", err)
		res.Error(api.ResultError)
		return
	}
	content := []byte("您的验证码为：" + code + ", 5分钟内有效，请及时操作。")
	ulits.SendMail(email, content)
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
func Register(c *gin.Context) {
	var (
		res      = api.NewResult(c)
		username = c.PostForm("username")
		password = c.PostForm("password")
		mobile   = c.PostForm("mobile")
		code     = c.PostForm("code")
	)
	if code == "" || username == "" || password == "" {
		res.Error(api.InvalidArgs)
		return
	}
	// 验证验证码是否正确
	verificationCode, err := models.Redis.Get(mobile).Result()
	if err != nil {
		log.Printf("Get Code Error:%v \n", err)
		res.Error(api.CodeExpire)
		return
	}
	if verificationCode != code {
		res.Error(api.CodeError)
		return
	}
	// 判断邮箱是否已存在
	var cnt int64
	err = models.DB.Where("mobile = ?", mobile).Model(new(models.User)).Count(&cnt).Error
	if err != nil {
		res.Error(api.DatabaseError)
		return
	}
	if cnt > 0 {
		res.Error(api.UserExist)
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
		res.Error(api.DatabaseError)
		return
	}

	// 生成 token
	token, err := ulits.GetToken(user.UUID, user.UserName, user.Role)
	if err != nil {
		res.Error(api.ResultError)
		return
	}
	res.Success(map[string]interface{}{
		"token": token,
	})
}

// Login
// @Tags 用户管理
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":"","msg":"success"}"
// @Router /login [post]
func Login(c *gin.Context) {
	res := api.NewResult(c)
	user := new(models.User)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		res.Error(api.InvalidArgs)
		return
	}
	// password = ulits.Md5(password)

	err := models.DB.Where("username = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res.Error(api.UserError)
			return
		}
		res.Error(api.DatabaseError)
		log.Printf("Models err: %s", err)
		return
	}

	token, err := ulits.GetToken(user.UUID, user.UserName, user.Role)
	if err != nil {
		res.Error(api.Unauthorized)
		log.Printf("GetToken err: %s", err)
		return
	}
	data := map[string]string{
		"token": token,
	}
	res.Success(data)
}

// GetUserInfo
// @Tags 用户管理
// @Summary 用户详情
// @Param authorization header string true "authorization"
// @Param id query string false "id"
// @Success 200 {string} json "{"code":"200","data":"","msg":"success"}"
// @Router /user/info [get]
func GetUserInfo(c *gin.Context) {
	res := api.NewResult(c)
	id := c.Query("id")
	if id == "" {

		res.Error(api.InvalidArgs)
		return
	}
	data := new(models.User)
	err := models.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		res.Error(api.NeedRedirect)
		return
	}
	res.Success(data)
}
