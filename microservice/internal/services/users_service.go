package services

type NewUserAccountRequest struct {
	UserName     string `json:"username"`
	HashPassword string `json:"password"`
	Email        string `json:"email"`
	AcceptTerms  bool   `json:"acceptterms"`
	IpAddress    string `json:"ipaddress"`
}
type NewUserAccountResponse struct {
	UserName     string `json:"username"`
	HashPassword string `json:"password"`
	Email        string `json:"email"`
	AcceptTerms  bool   `json:"acceptterms"`
	IpAddress    string `json:"ipaddress"`
	CreateAt     string `json:"createat"`
	UpdateAt     string `json:"updateat"`
	DeleteAt     string `json:"deleteat"`
}

type FindAccountReponse struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type UserService interface {
	NewUserAccount(newUserAccountRequest *NewUserAccountRequest) (*NewUserAccountResponse, error)
	FindAllAccount() ([]FindAccountReponse, error)
	FindAccount(string) (*FindAccountReponse, error)
}
