package models

import "database/sql"

// RegisterRequest represents a registration request payload.
type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

// RegisterResult represents the result of a registration attempt.
type RegisterResult struct {
	Error   int
	Message string
}

// LoginRequest represents a login request payload.
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// LoginResult represents the result of a login attempt.
type LoginResult struct {
	Error     int
	Message   string
	UID       sql.NullString
	UserName  sql.NullString
	Password  sql.NullString
	ExpiredAt sql.NullTime
}
