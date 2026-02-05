package models

import "database/sql"

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
