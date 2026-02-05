package models

import "database/sql"

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