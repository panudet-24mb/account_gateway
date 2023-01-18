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

type UserRoles struct {
	UserID   string `bson:"userid"`
	RoleID   string `bson:"roleid"`
	CreateAt string `bson:"createat"`
	UpdateAt string `bson:"updateat"`
	DeleteAt string `bson:"deleteat"`
}
