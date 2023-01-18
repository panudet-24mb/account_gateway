package services

import "account_gateway/internal/repository"

type userGroupService struct {
	userGroupRepository repository.UserGroupRepository
}

func NewUserGroupService(userGroupRepository repository.UserGroupRepository) UserGroupService {
	return userGroupService{userGroupRepository: userGroupRepository}
}

func (u userGroupService) NewGroup(newGroupRequest *NewGroupRequest) (*NewGroupResponse, error) {
	userGroup := repository.Group{
		GroupName:      newGroupRequest.GroupName,
		GroupDesc:      newGroupRequest.GroupDesc,
		GroupImgUrl:    newGroupRequest.GroupImgUrl,
		GroupIcon:      newGroupRequest.GroupIcon,
		GroupIconColor: newGroupRequest.GroupIconColor,
		GroupType:      newGroupRequest.GroupType,
	}
	_, err := u.userGroupRepository.Create(userGroup)
	if err != nil {
		return nil, err
	}

	return &NewGroupResponse{
		GroupName:      newGroupRequest.GroupName,
		GroupDesc:      newGroupRequest.GroupDesc,
		GroupImgUrl:    newGroupRequest.GroupImgUrl,
		GroupIcon:      newGroupRequest.GroupIcon,
		GroupIconColor: newGroupRequest.GroupIconColor,
		GroupType:      newGroupRequest.GroupType,
	}, nil
}
