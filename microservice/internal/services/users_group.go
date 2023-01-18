package services

import "account_gateway/internal/repository"

type userGroupService struct {
	userGroupRepository repository.UserGroupRepository
}

func NewUserGroupService(userGroupRepository repository.UserGroupRepository) UserGroupService {
	return userGroupService{userGroupRepository: userGroupRepository}
}

func (u userGroupService) NewUserGroup(newUserGroupRequest *NewUserGroupRequest) (*NewUserGroupResponse, error) {
	userGroup := repository.UserGroup{
		GroupName:      newUserGroupRequest.GroupName,
		GroupDesc:      newUserGroupRequest.GroupDesc,
		GroupImgUrl:    newUserGroupRequest.GroupImgUrl,
		GroupIcon:      newUserGroupRequest.GroupIcon,
		GroupIconColor: newUserGroupRequest.GroupIconColor,
		GroupType:      newUserGroupRequest.GroupType,
	}
	_, err := u.userGroupRepository.Create(userGroup)
	if err != nil {
		return nil, err
	}

	return &NewUserGroupResponse{
		GroupName:      newUserGroupRequest.GroupName,
		GroupDesc:      newUserGroupRequest.GroupDesc,
		GroupImgUrl:    newUserGroupRequest.GroupImgUrl,
		GroupIcon:      newUserGroupRequest.GroupIcon,
		GroupIconColor: newUserGroupRequest.GroupIconColor,
		GroupType:      newUserGroupRequest.GroupType,
	}, nil
}
