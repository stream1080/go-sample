package service

import (
	"demo/api"
	"demo/models"

	"github.com/gin-gonic/gin"
)

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
