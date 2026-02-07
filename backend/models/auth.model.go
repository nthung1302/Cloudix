package models

import "database/sql"

// Đăng nhập
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Status    int
	Message   string
	UID       sql.NullString
	UserName  sql.NullString
	Password  sql.NullString
	ExpiredAt sql.NullTime
}

// Quên mật khẩu
type ForgotPasswordResult struct {
	Status  int
	Message string
}

type ForgotPasswordRequest struct {
	UserName    string `json:"username"`
	NewPassword string `json:"new_password"`
}

// Đăng ký tài khoản
type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type RegisterResult struct {
	Status  int
	Message string
}
