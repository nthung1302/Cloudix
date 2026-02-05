package models

import "database/sql"

// Account represents a user account in the system.
type Account struct {
	UID       sql.NullString
	UserName  sql.NullString
	Password  sql.NullString
	FullName  sql.NullString
	Email     sql.NullString
	CreatedAt sql.NullTime
	ExpiredAt sql.NullTime
}

// ChangePasswordRequest represents a change password request payload.
type ChangePasswordRequest struct {
	NewPassword string `json:"new_password"`
}
// ChangePasswordResult represents the result of a change password attempt.
type ChangePasswordResult struct {
	Error   int
	Message string
}