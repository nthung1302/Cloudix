package models

type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type RegisterResult struct {
	Status  int
	Message string
}
