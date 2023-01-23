package services

type NewUserAccountRequest struct {
	UserName     string `json:"username"`
	HashPassword string `json:"password"`
	Email        string `json:"email"`
	AcceptTerms  bool   `json:"acceptterms"`
}
type NewUserAccountResponse struct {
	UserName     string `json:"username"`
	HashPassword string `json:"password"`
	Email        string `json:"email"`
	AcceptTerms  bool   `json:"acceptterms"`
	CreateAt     string `json:"createat"`
	UpdateAt     string `json:"updateat"`
	DeleteAt     string `json:"deleteat"`
}

type UpdateUserRequest struct {
	UserName     string `json:"username"`
	Email        string `json:"email"`
	HashPassword string `json:"password"`
	LastLogin    string `json:"lastlogin"`
	UpdateAt     string `json:"updateat"`
}

type UpdateUserResponse struct {
	UserName  string `json:"username"`
	Email     string `json:"email"`
	LastLogin string `json:"lastlogin"`
	UpdateAt  string `json:"updateat"`
}
type DeleteAccountRequest struct {
	UserName string `json:"username"`
	DeleteAt string `json:"deleteat"`
}

type DeleteAccountResponse struct {
	UserName string `json:"username"`
	DeleteAt string `json:"deleteat"`
}

type FindAccountRequest struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	AcceptTerms bool   `json:"acceptterms"`
	CreateAt    string `json:"createat"`
	UpdateAt    string `json:"updateat"`
	DeleteAt    string `json:"deleteat"`
}
type FindAccountReponse struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	AcceptTerms bool   `json:"acceptterms"`
	CreateAt    string `json:"createat"`
	UpdateAt    string `json:"updateat"`
	DeleteAt    string `json:"deleteat"`
}

type UserService interface {
	NewUserAccount(newUserAccountRequest *NewUserAccountRequest) (*NewUserAccountResponse, error)
	GetAccounts() ([]FindAccountReponse, error)
	GetAccountByID(id string) (*FindAccountReponse, error)
	UpdateAccounts(updateUser *UpdateUserRequest) (*UpdateUserResponse, error)
	UpdateAccount(updateUser *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteAccount(username string, deleteUser *DeleteAccountRequest) (*DeleteAccountResponse, error)
	FindAccount(findAccount FindAccountRequest) (*FindAccountReponse, error)
}
