package services

import (
	"backend/models"
	"backend/repositories"
)

// Thay đổi mật khẩu
func ChangePassword(UID, NewPassword string) (*models.ChangePasswordResult, error) {
	return repositories.ChangePassword(UID, NewPassword)
}

// Lấy thông tin tài khoản
func GetAccountInfo(uid string) (map[string]interface{}, error) {
	return repositories.GetAccountInfo(uid)
}

// Lấy quyền tài khoản
func GetAccountPermissions(uid string) ([]models.PermissionRow, error) {
	return repositories.GetAccountPermissions(uid)
}
