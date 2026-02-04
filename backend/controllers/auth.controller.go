package controllers

import (
	"backend/services"

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
func RegisterAccount(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "Dữ liệu không hợp lệ"})
		return
	}

	if err := services.Register(
		req.Username,
		req.Password,
		req.FullName,
		req.Email,
	); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Đăng ký thành công"})
}

/*	Đăng nhập tài khoản
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
func LoginAccount(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "Dữ liệu không hợp lệ"})
		return
	}

	token, err := services.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"access_token": token})
}
