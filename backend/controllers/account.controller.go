package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

// ChangePassword handles the password change request for an authenticated user.
func ChangePassword(c *gin.Context) {
	UID := c.GetString("UID")
	if UID == "" {
		utils.Logger.Error("Không tìm thấy UID trong context.")
		utils.Response(c, 401, 0, "Unauthorized.", nil)
		return
	}

	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Logger.Error("Bind JSON lỗi: ", err)
		utils.Response(c, 400, 0, "Dữ liệu không hợp lệ.", nil)
		return
	}

	result, err := services.ChangePassword(UID, req.NewPassword)
	if err != nil {
		utils.Logger.Error("Đổi mật khẩu lỗi: ", err)
		utils.Response(c, 400, 0, err.Error(), nil)
		return
	}

	utils.Response(c, 200, result.Status, result.Message, nil)
}

// GetAccountInfo retrieves the account information for an authenticated user.
func GetAccountInfo(c *gin.Context) {
	uidVal, ok := c.Get("UID")
	if !ok {
		utils.Logger.Error("Không tìm thấy UID trong context.")
		utils.Response(c, 401, 0, "Unauthorized.", nil)
		return
	}
	UID := uidVal.(string)

	// PROFILE
	profile, err := services.GetAccountInfo(UID)
	if err != nil {
		utils.Logger.Error(err)
		utils.Response(c, 400, 0, err.Error(), nil)
		return
	}

	// RAW PERMISSIONS
	rawPermissions, err := services.GetAccountPermissions(UID)
	if err != nil {
		utils.Logger.Error(err)
		utils.Response(c, 400, 0, err.Error(), nil)
		return
	}

	// GROUP PERMISSIONS
	groupMap := make(map[string][]models.PermissionItem)

	for _, p := range rawPermissions {
		groupName := p.GroupName.String

		groupMap[groupName] = append(groupMap[groupName], models.PermissionItem{
			PermissionName: p.PermissionName,
		})
	}

	// CONVERT MAP -> ARRAY
	var permissionGroups []models.PermissionGroup
	for groupName, perms := range groupMap {
		permissionGroups = append(permissionGroups, models.PermissionGroup{
			GroupName:   groupName,
			Permissions: perms,
		})
	}

	utils.Response(c, 200, 0, "Lấy thông tin tài khoản thành công.", gin.H{
		"profile":     profile,
		"permissions": permissionGroups,
	})
}

