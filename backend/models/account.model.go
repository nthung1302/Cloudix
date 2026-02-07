package models

import "database/sql"

// Thông tin tài khoản và quyền người dùng
type AccountProfileResult struct {
	Status    int
	Message   string
	UserName  sql.NullString
	FullName  sql.NullString
	Email     sql.NullString
	CreatedAt sql.NullTime
}

type PermissionItem struct {
	PermissionName string `json:"permission_name"`
}

type PermissionGroup struct {
	GroupName   string           `json:"group_name"`
	Permissions []PermissionItem `json:"permissions"`
}

type PermissionRow struct {
	GroupName      sql.NullString
	PermissionName string
}

// Lấy nhật ký tài khoản
type AuditLogresult struct {
	LogID     sql.NullInt64
	AccountID sql.NullString
	Action    sql.NullString
	TimeStamp sql.NullTime
	Details   sql.NullString
}


// Đổi mật khẩu
type ChangePasswordRequest struct {
	NewPassword string `json:"new_password"`
}

type ChangePasswordResult struct {
	Status  int
	Message string
}

