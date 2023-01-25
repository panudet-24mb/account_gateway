package services

type NewGroupRequest struct {
	GroupName      string `json:"groupname"`
	GroupDesc      string `json:"groupdesc"`
	GroupImgUrl    string `json:"groupimgurl"`
	GroupIcon      string `json:"groupicon"`
	GroupIconColor string `json:"groupiconcolor"`
	GroupType      string `json:"grouptype"`
}

type NewGroupResponse struct {
	GroupName      string `json:"groupname"`
	GroupCode      string `json:"groupcode"`
	GroupDesc      string `json:"groupdesc"`
	GroupImgUrl    string `json:"groupimgurl"`
	GroupIcon      string `json:"groupicon"`
	GroupIconColor string `json:"groupiconcolor"`
	GroupType      string `json:"grouptype"`
	GroupStatus    string `json:"groupstatus"`
	GroupActive    bool   `json:"groupactive"`
	CreateAt       string `json:"createat"`
	UpdateAt       string `json:"updateat"`
}

type UpdateGroupRequest struct {
	GroupName      string `json:"groupname"`
	GroupCode      string `json:"groupcode"`
	GroupDesc      string `json:"groupdesc"`
	GroupImgUrl    string `json:"groupimgurl"`
	GroupIcon      string `json:"groupicon"`
	GroupIconColor string `json:"groupiconcolor"`
	GroupType      string `json:"grouptype"`
	GroupStatus    string `json:"groupstatus"`
	GroupActive    bool   `json:"groupactive"`
	UpdateAt       string `json:"updateat"`
	DeleteAt       string `json:"deleteat"`
}

type UpdateGroupResponse struct {
	GroupName      string `json:"groupname"`
	GroupCode      string `json:"groupcode"`
	GroupDesc      string `json:"groupdesc"`
	GroupImgUrl    string `json:"groupimgurl"`
	GroupIcon      string `json:"groupicon"`
	GroupIconColor string `json:"groupiconcolor"`
	GroupType      string `json:"grouptype"`
	GroupStatus    string `json:"groupstatus"`
	GroupActive    bool   `json:"groupactive"`
	UpdateAt       string `json:"updateat"`
	DeleteAt       string `json:"deleteat"`
}

type FindGroupRequest struct {
	GroupName      string `json:"groupname"`
	GroupCode      string `json:"groupcode"`
	GroupDesc      string `json:"groupdesc"`
	GroupImgUrl    string `json:"groupimgurl"`
	GroupIcon      string `json:"groupicon"`
	GroupIconColor string `json:"groupiconcolor"`
	GroupType      string `json:"grouptype"`
	GroupStatus    string `json:"groupstatus"`
	GroupActive    bool   `json:"groupactive"`
	CreateAt       string `json:"createat"`
	UpdateAt       string `json:"updateat"`
}

type FindGroupResponse struct {
	GroupName      string `json:"groupname"`
	GroupCode      string `json:"groupcode"`
	GroupDesc      string `json:"groupdesc"`
	GroupImgUrl    string `json:"groupimgurl"`
	GroupIcon      string `json:"groupicon"`
	GroupIconColor string `json:"groupiconcolor"`
	GroupType      string `json:"grouptype"`
	GroupStatus    string `json:"groupstatus"`
	GroupActive    bool   `json:"groupactive"`
	CreateAt       string `json:"createat"`
	UpdateAt       string `json:"updateat"`
}

type UserGroupService interface {
	NewGroup(newUserGroupRequest *NewGroupRequest) (*NewGroupResponse, error)
	GetGroups() ([]NewGroupResponse, error)
	GetGroupByID(id string) (*NewGroupResponse, error)
	UpdateGroups(updateGroup *UpdateGroupRequest) (*UpdateGroupResponse, error)
	UpdateGroup(id string, updateGroup *UpdateGroupRequest) (*UpdateGroupResponse, error)
	DeleteGroup(groupName string, updateGroup *UpdateGroupRequest) (*UpdateGroupResponse, error)
	// FindGroup(findGroup FindGroupRequest) ([]FindGroupResponse, error)
}
