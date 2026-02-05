package services

import (
	"backend/repositories"
	"backend/models"
)

//	Đăng ký tài khoản
func Register(username, password, fullname, email string) (*models.RegisterResult, error) {
	return repositories.Register(username, password, fullname, email)
}

//	Đăng nhập tài khoản
func Login(username, password string) (*models.LoginResult, error) {
	return repositories.Login(username, password)
}

// Quên mật khẩu
func ForgotPassword(username, newPassword string) (*models.ForgotPasswordResult, error) {
	return repositories.ForgotPassword(username, newPassword)
}
