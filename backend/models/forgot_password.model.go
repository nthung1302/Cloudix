package models


type ForgotPasswordResult struct {
	Status  int
	Message string
}

type ForgotPasswordRequest struct {
	UserName    string `json:"username"`
	NewPassword string `json:"new_password"`
}
