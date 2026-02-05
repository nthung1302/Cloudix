package repositories

import (
	"backend/configs"
	"backend/utils"
	"backend/models"
	"database/sql"
)

func Register(Username, Password, FullName, Email string) (*models.RegisterResult, error) {
	row := configs.DB.QueryRow(
		`EXEC SP_CREATE_ACCOUNT @UserName=@u, @Password=@p,	@FullName=@f, @Email=@e`,
		sql.Named("u", Username),
		sql.Named("p", Password),
		sql.Named("f", FullName),
		sql.Named("e", Email),
	)

	var res models.RegisterResult
	if err := row.Scan(&res.Error, &res.Message); err != nil {
		utils.DebugLog("Error scanning register result", err.Error())
		return nil, err
	}

	return &res, nil
}

func Login(UserName string) (*models.LoginResult, error) {
	row := configs.DB.QueryRow(
		`EXEC SP_LOGIN_ACCOUNT @UserName=@u`,
		sql.Named("u", UserName),
	)

	var res models.LoginResult
	err := row.Scan( &res.Error, &res.Message, &res.UID, &res.UserName, &res.Password, &res.ExpiredAt,)

	if err != nil {
		utils.DebugLog("Error scanning login result", err.Error())
		return nil, err
	}

	return &res, nil
}
