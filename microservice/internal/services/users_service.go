package services

type NewUserAccountRequest struct {
	UserName     string `json:"username"`
	HashPassword string `json:"password"`
	Email        string `json:"email"`
	AcceptTerms  bool   `json:"acceptterms"`
	IpAddress    string `json:"ipaddress"`
}
type NewUserAccountResponse struct {
	UserName string `json:"username"`
}

type UserService interface {
	NewUserAccount(newUserAccountRequest *NewUserAccountRequest) (*NewUserAccountResponse, error)
}
