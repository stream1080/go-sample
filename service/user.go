package service

import (
	"demo/api"
	"demo/models"
	"demo/ulits"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":"","msg":""}"
// @Router /user-login [post]
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
// @Tags 公共方法
// @Summary 用户详情
// @Param id query string false "id"
// @Success 200 {string} json "{"code":"200","data":"","msg":""}"
// @Router /user-info [get]
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
