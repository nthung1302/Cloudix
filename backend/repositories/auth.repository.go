package repositories

import (
	"database/sql"
	"backend/configs"
)

type ProcResult struct {
	Result  int
	Message string
}

func CreateUserByProc(username, password string) (*ProcResult, error) {
	row := configs.DB.QueryRow(
	"EXEC LOGIN @username = @Username, @password = @Password",
		sql.Named("Username", username),
		sql.Named("Password", password),
	)

	var res ProcResult
	err := row.Scan(&res.Result, &res.Message)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
