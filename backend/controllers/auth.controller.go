package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

// Register handles the registration request.
func Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Lỗi khi nhận dữ liệu đăng ký: ", err)
		utils.Response(c, 400, "Dữ liệu không hợp lệ.", nil)
		return
	}

	result, err := services.Register(req.UserName, req.Password, req.FullName, req.Email)
	if err != nil {
		utils.Response(c, 400, err.Error(), nil)
		return
	}

	utils.Response(c, 200, result.Message, nil)
}

// Login handles the login request.
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, 400, "Dữ liệu không hợp lệ.", nil)
		return
	}

	user, err := services.Login(req.UserName, req.Password)
	if err != nil {
		utils.Logger.Error(err)
		utils.Response(c, 400, err.Error(), nil)
		return
	}

	token, err := utils.GenerateAccessToken(
		user.UID.String,
		user.UserName.String,
	)
	if err != nil {
		utils.Logger.Error(err)
		utils.Response(c, 500, "Lỗi hệ thống.", nil)
		return
	}

	utils.Response(c, 200, "Đăng nhập thành công.", gin.H{
		"access_token": token,
	})
}

// ForgotPassword handles the forgot password request.
func ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Lỗi khi nhận dữ liệu quên mật khẩu: ", err)
		utils.Response(c, 400, "Dữ liệu không hợp lệ.", nil)
		return
	}

	result, err := services.ForgotPassword(req.UserName, req.NewPassword)
	if err != nil {
		utils.Logger.Error("Lỗi khi xử lý quên mật khẩu: ", err)
		utils.Response(c, 400, err.Error(), nil)
		return
	}

	utils.Response(c, 200, result.Message, nil)
}
