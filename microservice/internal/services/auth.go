package services

type LoginDefaultForm struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ResponseLogin struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type AuthService interface {
	DefaultLogin(LoginDefaultForm *LoginDefaultForm) (*ResponseLogin, error)
}
