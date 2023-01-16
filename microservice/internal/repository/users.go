package repository

type User struct {
	UserID         uint
	UserName       string
	HashPassword   string
	Email          string
	EmailConfirm   bool
	UserStatus     string
	LoginAttempt   int
	AcceptTerms    bool
	VerifyIdentity bool
	LastLogin      string
	IpAddress      string
	RegisterAt     string
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(int) (*User, error)
	Create(User) (*User, error)
	Update(User) (*User, error)
	UpdateOne(User) (*User, error)
	DeleteOne(int) error
	FindUser(User) (*User, error)
}
