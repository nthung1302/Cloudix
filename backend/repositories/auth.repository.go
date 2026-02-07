package repositories

import (
	"backend/configs"
	"backend/models"
	"backend/utils"
	"database/sql"
	"errors"
	"time"
)

// Đăng ký tài khoản
func Register(Username, Password, FullName, Email string) (*models.RegisterResult, error) {
	hash, err := utils.HashPassword(Password)
	if err != nil {
		utils.Logger.Error("Lỗi khi mã hóa mật khẩu: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	row := configs.DB.QueryRow(
		`EXEC SP_CREATE_ACCOUNT @UserName=@u, @Password=@p,	@FullName=@f, @Email=@e`,
		sql.Named("u", Username),
		sql.Named("p", hash),
		sql.Named("f", FullName),
		sql.Named("e", Email),
	)

	var res models.RegisterResult
	if err := row.Scan(&res.Status, &res.Message); err != nil {
		utils.Logger.Error("Scan SP_CREATE_ACCOUNT lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status < 0 {
		utils.Logger.Error("SP_CREATE_ACCOUNT trả về status < 0: " + res.Message)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status == 0 {
		return nil, errors.New(res.Message)
	}
	return &res, nil
}

// Đăng nhập tài khoản
func Login(UserName, Password string) (*models.LoginResult, error) {
	row := configs.DB.QueryRow(
		`EXEC SP_LOGIN_ACCOUNT @UserName=@userName`,
		sql.Named("userName", UserName),
	)

	var res models.LoginResult
	err := row.Scan(&res.Status, &res.Message, &res.UID, &res.UserName, &res.Password, &res.ExpiredAt)
	if err != nil {
		utils.Logger.Error("Scan SP_LOGIN_ACCOUNT lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống")
	}

	if res.Status < 0 {
		utils.Logger.Error("SP_LOGIN_ACCOUNT trả về status < 0: " + res.Message)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status == 0 {
		return nil, errors.New(res.Message)
	}

	if !utils.CheckPassword(res.Password.String, Password) {
		return nil, errors.New("Sai mật khẩu.")
	}

	if res.ExpiredAt.Valid && res.ExpiredAt.Time.Before(time.Now()) {
		return nil, errors.New("Tài khoản đã hết hạn.")
	}

	return &res, nil
}

// Quên mật khẩu
func ForgotPassword(Username, NewPassword string) (*models.ForgotPasswordResult, error) {
	hash, err := utils.HashPassword(NewPassword)
	if err != nil {
		utils.Logger.Error("Lỗi khi mã hóa mật khẩu mới: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	row := configs.DB.QueryRow(
		`EXEC SP_UPDATE_PASSWORD_BY_USERNAME @UserName=@userName, @NewPassword=@newPassword`,
		sql.Named("userName", Username),
		sql.Named("newPassword", hash),
	)

	var res models.ForgotPasswordResult
	if err := row.Scan(&res.Status, &res.Message); err != nil {
		utils.Logger.Error("Scan SP_UPDATE_PASSWORD_BY_USERNAME lỗi: ", err)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status < 0 {
		utils.Logger.Error("SP_UPDATE_PASSWORD_BY_USERNAME trả về status < 0: " + res.Message)
		return nil, errors.New("Lỗi hệ thống.")
	}

	if res.Status == 0 {
		return nil, errors.New(res.Message)
	}

	return &res, nil
}