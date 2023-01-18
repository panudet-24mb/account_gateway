package repository

type UserGroup struct {
	GroupName string `bson:"groupname"`
	GroupDesc string `bson:"groupdesc"`
	GroupType string `bson:"grouptype"`
	CreateAt  string `bson:"createat"`
	UpdateAt  string `bson:"updateat"`
	DeleteAt  string `bson:"deleteat"`
}
