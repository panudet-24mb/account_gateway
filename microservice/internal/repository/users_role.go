package repository

type Roles struct {
	RoleName      string `bson:"rolename"`
	RoleCode      string `bson:"rolecode"`
	RoleDesc      string `bson:"roledesc"`
	RoleImgUrl    string `bson:"roleimgurl"`
	RoleIcon      string `bson:"roleicon"`
	RoleIconColor string `bson:"roleiconcolor"`
	RoleType      string `bson:"roletype"`
	RoleStatus    string `bson:"rolestatus"`
	RoleActive    bool   `bson:"roleactive"`
	CreateAt      string `bson:"createat"`
	UpdateAt      string `bson:"updateat"`
	DeleteAt      string `bson:"deleteat"`
}

type UsersRole struct {
	UserID   string `bson:"userid"`
	RoleID   string `bson:"roleid"`
	CreateAt string `bson:"createat"`
	UpdateAt string `bson:"updateat"`
	DeleteAt string `bson:"deleteat"`
}

type RoleRepository interface {
	Create(Roles) (*Roles, error)
	GetAll() ([]Roles, error)
	GetByID(string) (*Roles, error)
	Update(Roles) (*Roles, error)
	UpdateOne(string, Roles) (*Roles, error)
	DeleteOne(string, Roles) (*Roles, error)
}

type UserRoleRepository interface {
	Create(UsersRole) (*UsersRole, error)
	GetAll() ([]UsersRole, error)
	GetByID(id string) (*UsersRole, error)
	Update(UsersRole) (*UsersRole, error)
	UpdateOne(string, UsersRole) (*UsersRole, error)
	DeleteOne(string, UsersRole) (*UsersRole, error)
}
