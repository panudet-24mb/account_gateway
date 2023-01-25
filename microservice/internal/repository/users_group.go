package repository

type Group struct {
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
	Create(Group Group) (*Group, error)
	GetAll() ([]Group, error)
	GetByID(id string) (*Group, error)
	Update(Group Group) (*Group, error)
	UpdateOne(id string, Group Group) (*Group, error)
	DeleteOne(groupName string, Group Group) (*Group, error)
	// FindGroup(g Group) ([]Group, error)
}
