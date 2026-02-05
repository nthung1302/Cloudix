package repositories

import (
	"backend/configs"
	"backend/models"
	"backend/utils"
	"database/sql"
)

func ChangePassword(UID, NewPassword string) (*models.ChangePasswordResult, error) {

	row := configs.DB.QueryRow(
		`EXEC SP_CHANGE_PASSWORD @UID=@u, @NewPassword=@n`,
		sql.Named("u", UID),
		sql.Named("n", NewPassword),
	)
	var res models.ChangePasswordResult
	if err := row.Scan(&res.Error, &res.Message); err != nil {
		utils.DebugLog("Error scanning change password result", err.Error())
		return nil, err
	}
	return &res, nil
}
