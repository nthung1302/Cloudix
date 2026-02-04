package services

import (
	"backend/repositories"
	"backend/utils"
	"errors"
	"time"
)

//========================================================================
//	Đăng ký tài khoản
//========================================================================
func Register(username, password, fullName, email string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	result, err := repositories.RegisterAccountRepo(
		username,
		hash,
		fullName,
		email,
	)
	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New(result.Message)
	}

	return nil
}

//========================================================================
//	Đăng nhập tài khoản
//========================================================================
func Login(username, password string) (string, error) {
	result, err := repositories.LoginAccountRepo(username)
	if err != nil {
		return "", err
	}

	if result.Status == 0 {
		return "", errors.New(result.Message)
	}

	if !utils.CheckPassword(result.PasswordHash.String, password) {
		return "", errors.New("Mật khẩu không đúng")
	}

	if result.ExpiredAt.Valid && result.ExpiredAt.Time.Before(time.Now()) {
		return "", errors.New("Tài khoản đã hết hạn")
	}

	token, err := utils.GenerateAccessToken(
		int(result.UserID.Int64),
		result.Username.String,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
