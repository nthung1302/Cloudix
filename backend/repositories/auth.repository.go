package repositories

import (
	"backend/common"
	"backend/configs"
	"database/sql"
)

func RegisterAccountRepo(username, password, fullName, email string) (*common.ProcResult, error) {
	row := configs.DB.QueryRow(
		`EXEC SP_CREATE_ACCOUNT @username=@u, @password=@p,	@full_name=@f, @email=@e`,
		sql.Named("u", username),
		sql.Named("p", password),
		sql.Named("f", fullName),
		sql.Named("e", email),
	)

	var res common.ProcResult
	if err := row.Scan(&res.Status, &res.Message); err != nil {
		return nil, err
	}

	return &res, nil
}

type LoginResult struct {
	Status       int
	Message      string
	UserID       sql.NullInt64
	Username     sql.NullString
	PasswordHash sql.NullString
	ExpiredAt    sql.NullTime
}

func LoginAccountRepo(login string) (*LoginResult, error) {
	row := configs.DB.QueryRow(
		`EXEC SP_LOGIN_ACCOUNT @username=@Username`,
		sql.Named("Username", login),
	)

	var res LoginResult
	err := row.Scan(
	&res.Status,
	&res.Message,
	&res.UserID,
	&res.Username,
	&res.PasswordHash,
	&res.ExpiredAt,
	)


	if err != nil {
		return nil, err // system error
	}

	return &res, nil
}
