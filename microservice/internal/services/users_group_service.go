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
}

type UserGroupService interface {
	NewGroup(newUserGroupRequest *NewGroupRequest) (*NewGroupResponse, error)
}
