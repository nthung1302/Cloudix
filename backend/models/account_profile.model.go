package models

import "database/sql"

type AccountProfileResult struct {
	Status    int
	Message   string
	UserName  sql.NullString
	FullName  sql.NullString
	Email     sql.NullString
	CreatedAt sql.NullTime
}