package repositories

import (
	"backend/configs"
	"backend/models"
	"backend/utils"
	"database/sql"
	"errors"
)

// Thay đổi mật khẩu
func ChangePassword(UID, NewPassword string) (*models.ChangePasswordResult, error) {
	passwordHash, err := utils.HashPassword(NewPassword)
	if err != nil {
		utils.Logger.Error("Hash password lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	row := configs.DB.QueryRow(
		`EXEC SP_UPDATE_PASSWORD_BY_UID @UID=@uid, @NewPassword=@newPassword`,
		sql.Named("uid", UID),
		sql.Named("newPassword", passwordHash),
	)

	var res models.ChangePasswordResult
	if err := row.Scan(&res.Status, &res.Message); err != nil {
		utils.Logger.Error("Scan SP_UPDATE_PASSWORD_BY_UID lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status == 0 {
		return nil, errors.New(res.Message)
	}

	if res.Status < 0 {
		utils.Logger.Error("SP_UPDATE_PASSWORD_BY_UID trả về status < 0: " + res.Message)
		return nil, errors.New("Lỗi hệ thống.")
	}

	return &res, nil
}



// Lấy thông tin tài khoản
func GetAccountInfo(uid string) (map[string]interface{}, error) {
	row := configs.DB.QueryRow(
		`EXEC SP_GET_INFO_BY_UID @UID=@uid`,
		sql.Named("uid", uid),
	)

	var res models.AccountProfileResult
	if err := row.Scan(&res.Status, &res.Message, &res.UserName, &res.FullName, &res.Email, &res.CreatedAt); err != nil {
		utils.Logger.Error("Scan SP_GET_INFO_BY_UID lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status < 0 {
		utils.Logger.Error("SP_GET_INFO_BY_UID trả về status < 0: " + res.Message)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status == 0 {
		return nil, errors.New(res.Message)
	}

	return map[string]interface{}{
		"user_name":  res.UserName.String,
		"full_name":  res.FullName.String,
		"email":      res.Email.String,
		"created_at": res.CreatedAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}


// Lấy quyền tài khoản
func GetAccountPermissions(uid string) ([]models.PermissionRow, error) {
	rows, err := configs.DB.Query(
		`EXEC SP_GET_PERMISSION_BY_UID @UID=@u`,
		sql.Named("u", uid),
	)
	if err != nil {
		utils.Logger.Error("Query SP_GET_PERMISSION_BY_UID lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}
	defer rows.Close()

	var result []models.PermissionRow
	for rows.Next() {
		var r models.PermissionRow
		if err := rows.Scan(&r.GroupName, &r.PermissionName); err != nil {
			utils.Logger.Error("Scan SP_GET_PERMISSION_BY_UID lỗi: ", err)
			return nil, errors.New("Lỗi hệ thống.")
		}
		result = append(result, r)
	}
	return result, nil
}
