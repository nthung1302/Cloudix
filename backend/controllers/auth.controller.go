package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

/*
Đăng ký tài khoản mới
@Sumary đăng ký tài khoản mới
@Description Đăng ký tài khoản mới với username, password, full_name và email
@Tags Authentication
@Accept json
@Produce json
@Param dữ liệu body body object{ username=string, password=string, full_name=string, email=string} true "Dữ liệu đăng ký"
@Success 200 {object} object{message=string} đang ký thành công
@Failure 400 {object} object{message=string} lỗi đăng ký
@Router /api/authentication/register [post]
*/
func Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DebugLog(c, 400, -1, "Data not valid.", err)
		return
	}

	if err := services.Register(
		req.UserName,
		req.Password,
		req.FullName,
		req.Email,
	); err != nil {
		utils.DebugLogRes(c, 400, -2, err)
		return
	}

	utils.DebugLogRes(c, 200, 0, "Registration successful.", nil)
}	

/*
Đăng nhập tài khoản
@Sumary Đăng nhập tài khoản
@Description Đăng nhập tài khoản với username và password
@Tags Authentication
@Accept json
@Produce json
@Param dữ liệu body body object{ username=string, password=string} true "Dữ liệu đăng nhập"
@Success 200 {object} object{access_token=string} đăng nhập thành công
@Failure 400 {object} object{message=string} lỗi đăng nhập
@Router /api/authentication/login [post]
*/
func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.DebugLogRes(c, 400, -1, err)
		return
	}

	token, err := services.Login(req.UserName, req.Password)
	if err != nil {
		utils.DebugLogRes(c, 400, -2, err)
		return
	}
	c.JSON(200, gin.H{"access_token": token})
}
