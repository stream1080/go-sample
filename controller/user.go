package controller

import (
	"time"

	"go-sample/global"
	"go-sample/models"
	"go-sample/pkg/encrypt"
	"go-sample/pkg/jwt"
	"go-sample/pkg/response"
	"go-sample/pkg/ulits"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserApi struct {
}

type EmailParam struct {
	Email string `json:"email" binding:"required"`
}

func (u *UserApi) SendCode(c *gin.Context) {

	var param EmailParam
	if err := c.ShouldBindJSON(&param); err != nil {
		zap.S().Error("<UserApi.SendCode> c.ShouldBindJSON() failed with ", err)
		response.Error(c, response.InvalidArgs)
		return
	}

	code := ulits.GetCode()
	_, err := global.RDB.Set(param.Email, code, 5*60*time.Minute).Result()
	if err != nil {
		zap.S().Errorf("Set Code Error:%v \n", err)
		response.Error(c, response.ServerError)
		return
	}
	content := []byte("您的验证码为：" + code + ", 5分钟内有效, 请及时操作。")
	ulits.SendMail(param.Email, content)

	response.Success(c, nil)
}

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	SmsCode  string `json:"sms_code" binding:"required"`
}

func (u *UserApi) Register(c *gin.Context) {

	var form RegisterForm
	if err := c.ShouldBindJSON(&form); err != nil {
		zap.S().Error("<UserApi.Register> c.ShouldBindJSON() failed with ", err)
		response.Error(c, response.InvalidArgs)
		return
	}

	// 验证验证码是否正确
	verificationCode, err := global.RDB.Get(form.Mobile).Result()
	if err != nil {
		zap.S().Errorf("Get Code Error:%v \n", err)
		response.Error(c, response.CodeExpire)
		return
	}

	if verificationCode != form.SmsCode {
		response.Error(c, response.CodeError)
		return
	}
	// 判断邮箱是否已存在
	var cnt int64
	err = global.DB.Where("mobile = ?", form.Mobile).Model(new(models.User)).Count(&cnt).Error
	if err != nil {
		response.Error(c, response.ServerError)
		return
	}
	if cnt > 0 {
		response.Error(c, response.UserExist)
		return
	}

	// 数据的插入
	uuid := ulits.GetUUID()
	user := &models.User{
		UUID:     uuid,
		UserName: form.Username,
		Password: encrypt.Md5(form.Password),
		Mobile:   form.Mobile,
	}
	err = global.DB.Create(user).Error
	if err != nil {
		response.Error(c, response.ServerError)
		return
	}

	// 生成 token
	userClaims := jwt.NewClaims(user.UUID, user.UserName, user.Role)
	token, err := jwt.NewToken(userClaims, "")
	if err != nil {
		response.Error(c, response.ServerError)
		return
	}

	data := map[string]string{
		"token": token,
	}

	response.Success(c, data)
}

func (u *UserApi) Login(c *gin.Context) {
	user := new(models.User)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		response.Error(c, response.InvalidArgs)
		return
	}
	// password = ulits.Md5(password)

	if err := global.DB.Where("username = ? and password = ?", username, password).First(&user).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			response.Error(c, response.InvalidArgs)
			return
		}

		zap.S().Errorf("Models err: %s", err)
		response.Error(c, response.ServerError)

		return
	}

	userClaims := jwt.NewClaims(user.UUID, user.UserName, user.Role)
	token, err := jwt.NewToken(userClaims, "")
	if err != nil {
		response.Error(c, response.Unauthorized)
		zap.S().Errorf("GetToken err: %s", err)
		return
	}

	data := map[string]string{
		"token": token,
	}

	response.Success(c, data)
}

// UserInfo
// @Tags 用户管理
// @Summary 用户详情
// @Param authorization header string true "authorization"
// @Param id query string false "id"
// @Success 200 {string} json "{"code":"200","data":"","msg":"success"}"
// @Router /user/info [get]
func (u *UserApi) UserInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.Error(c, response.InvalidArgs)
		return
	}
	data := new(models.User)
	err := global.DB.Where("id = ?", id).Find(&data).Error
	if err != nil {
		response.Error(c, response.InvalidArgs)
		return
	}

	response.Success(c, data)
}
