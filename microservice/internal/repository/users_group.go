package repository

type UserGroup struct {
	GroupName      string `bson:"groupname"`
	GroupCode      string `bson:"groupcode"`
	GroupDesc      string `bson:"groupdesc"`
	GroupImgUrl    string `bson:"groupimgurl"`
	GroupIcon      string `bson:"groupicon"`
	GroupIconColor string `bson:"groupiconcolor"`
	GroupType      string `bson:"grouptype"`
	GroupStatus    string `bson:"groupstatus"`
	GroupActive    bool   `bson:"groupactive"`
	CreateAt       string `bson:"createat"`
	UpdateAt       string `bson:"updateat"`
	DeleteAt       string `bson:"deleteat"`
}

type UserGroupRepository interface {
	Create(UserGroup UserGroup) (*UserGroup, error)
}