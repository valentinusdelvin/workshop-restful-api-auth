package model

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string
	Password string
}

var (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)
