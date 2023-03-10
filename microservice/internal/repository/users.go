package repository

type User struct {
	UserName       string `bson:"username"`
	HashPassword   string `bson:"hashpassword"`
	Email          string `bson:"email"`
	EmailConfirm   bool   `bson:"emailconfirm"`
	UserStatus     string `bson:"userstatus"`
	LoginAttempt   int    `bson:"loginattemp"`
	AcceptTerms    bool   `bson:"acceptterms"`
	VerifyIdentity bool   `bson:"verifyidentity"`
	LastLogin      string `bson:"lastlogin"`
	CreateAt       string `bson:"createat"`
	UpdateAt       string `bson:"updateat"`
	DeleteAt       string `bson:"deleteat"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(string) (*User, error)
	Create(User) (*User, error)
	Update(User) (*User, error)
	UpdateOne(string, User) (*User, error)
	DeleteOne(string, User) (*User, error)
	FindUser(User) (*User, error)
}
