package services

import (
	"backend/repositories"
	"backend/utils"
	"errors"
	"time"
)

// ========================================================================
//	Đăng ký tài khoản
// ========================================================================
func Register(UserName, Password, FullName, Email string) error {
	PasswordHash, err := utils.HashPassword(Password)
	if err != nil {
		utils.DebugLog("Error hashing Password during registration", err.Error())
		return err
	}

	result, err := repositories.Register(UserName, PasswordHash, FullName, Email)
	if err != nil {
		utils.DebugLog("Error during registration", err.Error())
		return err
	}

	if result.Error != 0 {
		utils.DebugLog("Registration error", result.Message)
		return errors.New(result.Message)
	}

	return nil
}

// ========================================================================
//	Đăng nhập tài khoản
// ========================================================================
func Login(UserName, Password string) (string, error) {
	result, err := repositories.Login(UserName)
	if err != nil {
		utils.DebugLog("Error during login", err.Error())
		return "", err
	}

	if result.Error != 0 {
		utils.DebugLog("Login error", result.Message)
		return "", errors.New(result.Message)
	}

	if !utils.CheckPassword(result.Password.String, Password) {
		return "", errors.New("Password is incorrect.")
	}

	if result.ExpiredAt.Valid && result.ExpiredAt.Time.Before(time.Now()) {
		return "", errors.New("Account has expired.")
	}

	token, err := utils.GenerateAccessToken(
		result.UID.String,
		result.UserName.String,
	)
	if err != nil {
		utils.DebugLog("Error generating access token", err.Error())
		return "", err
	}

	return token, nil
}
