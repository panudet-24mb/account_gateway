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
	IpAddress      string `bson:"ipaddress"`
	RegisterAt     string `bson:"registerat"`
}

type UserRepository interface {
	// GetAll() ([]User, error)
	// GetById(int) (*User, error)
	Create(User) (*User, error)
	// Update(User) (*User, error)
	// UpdateOne(User) (*User, error)
	// DeleteOne(int) error
	// FindUser(User) (*User, error)
}
