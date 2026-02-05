package models

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password"`
}

type ChangePasswordResult struct {
	Status  int
	Message string
}