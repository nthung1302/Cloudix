package services

import (
	"backend/repositories"
	"backend/utils"
	"errors"
)

func ChangePassword(UID, NewPassword string) error {

	PasswordHash, err := utils.HashPassword(NewPassword)
	if err != nil {
		return err
	}
	result, err := repositories.ChangePassword(UID, PasswordHash)
	if err != nil {
		return err
	}
	if result.Error == 0 {
		return errors.New(result.Message)
	}
	return nil
}