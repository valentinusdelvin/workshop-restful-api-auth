package model

type UserRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string
	Password string
}

type UserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

var (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)
