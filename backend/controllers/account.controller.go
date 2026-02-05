package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

/* Đổi mật khẩu
@Sumary Đổi mật khẩu
@Description Đổi mật khẩu cho tài khoản đã đăng nhập
@Tags Account
@Accept json
@Produce json
@Param dữ liệu body body object{ new_password=string} true "Mật khẩu mới"
@Success 200 {object} object{message=string} đổi mật khẩu thành công
@Failure 400 {object} object{message=string} lỗi đổi mật khẩu
@Router /api/account/change-password [post]
*/

func ChangePassword(c *gin.Context) {
	var req models.ChangePasswordRequest
	UID := c.MustGet("UID").(string)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error":   -1,
			"message": "Data not valid.",
		})
		return
	}

	if err := services.ChangePassword(
		UID,
		req.NewPassword,
	); 
	err != nil {
		utils.DebugLogRes(c, 400, -2, "Change password failed.", err)
		return
	}

	utils.DebugLogRes(c, 200, 0, "Password changed successfully.", nil)
}

func GetAccountInfo(c *gin.Context) {
	UID := c.MustGet("UID").(string)
	UserName := c.MustGet("UserName").(string)
	c.JSON(200, gin.H{
		"user_id":  UID,
		"username": UserName,
	})
}
