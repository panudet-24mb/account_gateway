package services

import (
	"account_services/internal/repository"
	"log"
	"strconv"
	"time"
)

type userGroupService struct {
	userGroupRepository repository.UserGroupRepository
}

func NewUserGroupService(userGroupRepository repository.UserGroupRepository) UserGroupService {
	return userGroupService{userGroupRepository: userGroupRepository}
}

func (s userGroupService) NewGroup(newGroupRequest *NewGroupRequest) (*NewGroupResponse, error) {
	userGroup := repository.Group{
		GroupName:      newGroupRequest.GroupName,
		GroupDesc:      newGroupRequest.GroupDesc,
		GroupImgUrl:    newGroupRequest.GroupImgUrl,
		GroupIcon:      newGroupRequest.GroupIcon,
		GroupIconColor: newGroupRequest.GroupIconColor,
		GroupType:      newGroupRequest.GroupType,
		CreateAt:       strconv.Itoa(int(time.Now().Unix())),
	}
	data, err := s.userGroupRepository.Create(userGroup)
	if err != nil {
		return nil, err
	}

	userGroupResponse := &NewGroupResponse{
		GroupName:      data.GroupName,
		GroupDesc:      data.GroupDesc,
		GroupImgUrl:    data.GroupImgUrl,
		GroupIcon:      data.GroupIcon,
		GroupIconColor: data.GroupIconColor,
		GroupType:      data.GroupType,
		CreateAt:       data.CreateAt,
	}

	return userGroupResponse, nil
}

func (s userGroupService) GetGroups() ([]NewGroupResponse, error) {
	userGroups, err := s.userGroupRepository.GetAll()
	if err != nil {
		return nil, err
	}
	userGroupResponses := []NewGroupResponse{}
	for _, userGroup := range userGroups {
		userGroupResponse := NewGroupResponse{
			GroupName:      userGroup.GroupName,
			GroupDesc:      userGroup.GroupDesc,
			GroupImgUrl:    userGroup.GroupImgUrl,
			GroupIcon:      userGroup.GroupIcon,
			GroupIconColor: userGroup.GroupIconColor,
			GroupType:      userGroup.GroupType,
			CreateAt:       userGroup.CreateAt,
			UpdateAt:       userGroup.UpdateAt,
		}
		userGroupResponses = append(userGroupResponses, userGroupResponse)
	}
	return userGroupResponses, nil
}

func (s userGroupService) GetGroupByID(id string) (*NewGroupResponse, error) {
	userGroup, err := s.userGroupRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	userGroupResponse := NewGroupResponse{
		GroupName:      userGroup.GroupName,
		GroupDesc:      userGroup.GroupDesc,
		GroupImgUrl:    userGroup.GroupImgUrl,
		GroupIcon:      userGroup.GroupIcon,
		GroupIconColor: userGroup.GroupIconColor,
		GroupType:      userGroup.GroupType,
		CreateAt:       userGroup.CreateAt,
		UpdateAt:       userGroup.UpdateAt,
	}
	return &userGroupResponse, nil
}

func (s userGroupService) UpdateGroups(updateGroup *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	groupData := repository.Group{
		GroupName:      updateGroup.GroupName,
		GroupDesc:      updateGroup.GroupDesc,
		GroupImgUrl:    updateGroup.GroupImgUrl,
		GroupIcon:      updateGroup.GroupIcon,
		GroupIconColor: updateGroup.GroupIconColor,
		GroupType:      updateGroup.GroupType,
		GroupStatus:    updateGroup.GroupStatus,
		UpdateAt:       strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.userGroupRepository.Update(groupData)
	if err != nil {
		return nil, err
	}
	groupUpdate := &UpdateGroupResponse{
		GroupName:      dataUpdate.GroupName,
		GroupDesc:      dataUpdate.GroupDesc,
		GroupImgUrl:    dataUpdate.GroupImgUrl,
		GroupIcon:      dataUpdate.GroupIcon,
		GroupIconColor: dataUpdate.GroupIconColor,
		GroupType:      dataUpdate.GroupType,
		GroupStatus:    dataUpdate.GroupStatus,
		UpdateAt:       dataUpdate.UpdateAt,
	}
	return groupUpdate, nil
}

func (s userGroupService) UpdateGroup(id string, updateGroup *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	groupData := repository.Group{
		GroupName:      updateGroup.GroupName,
		GroupDesc:      updateGroup.GroupDesc,
		GroupImgUrl:    updateGroup.GroupImgUrl,
		GroupIcon:      updateGroup.GroupIcon,
		GroupIconColor: updateGroup.GroupIconColor,
		GroupType:      updateGroup.GroupType,
		GroupStatus:    updateGroup.GroupStatus,
		UpdateAt:       strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.userGroupRepository.UpdateOne(id, groupData)
	if err != nil {
		return nil, err
	}
	groupUpdate := &UpdateGroupResponse{
		GroupName:      dataUpdate.GroupName,
		GroupDesc:      dataUpdate.GroupDesc,
		GroupImgUrl:    dataUpdate.GroupImgUrl,
		GroupIcon:      dataUpdate.GroupIcon,
		GroupIconColor: dataUpdate.GroupIconColor,
		GroupType:      dataUpdate.GroupType,
		GroupStatus:    dataUpdate.GroupStatus,
		UpdateAt:       dataUpdate.UpdateAt,
	}
	return groupUpdate, nil
}

func (s userGroupService) DeleteGroup(groupName string, updateGroup *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	groupData := repository.Group{
		GroupName: updateGroup.GroupName,
		DeleteAt:  strconv.Itoa(int(time.Now().Unix())),
	}
	dataDelete, err := s.userGroupRepository.DeleteOne(groupName, groupData)
	if err != nil {
		log.Fatal(err)
	}
	groupDelete := &UpdateGroupResponse{
		GroupName: dataDelete.GroupName,
		DeleteAt:  dataDelete.DeleteAt,
	}
	return groupDelete, nil
}

// func (u userGroupService) FindGroup(findGroup FindGroupRequest) ([]FindGroupResponse, error) {
// 	groupData := repository.Group{
// 		GroupName:   findGroup.GroupName,
// 		GroupCode:   findGroup.GroupCode,
// 		GroupDesc:   findGroup.GroupDesc,
// 		GroupType:   findGroup.GroupType,
// 		GroupStatus: findGroup.GroupStatus,
// 		CreateAt:    findGroup.CreateAt,
// 		UpdateAt:    findGroup.UpdateAt,
// 	}
// 	groups, err := u.userGroupRepository.FindGroup(groupData)
// 	if err != nil {
// 		return nil, err
// 	}
// 	groupResponses := []FindGroupResponse{}
// 	for _, group := range groups {
// 		groupResponse := FindGroupResponse{
// 			GroupName:      group.GroupName,
// 			GroupCode:      group.GroupCode,
// 			GroupDesc:      group.GroupDesc,
// 			GroupImgUrl:    group.GroupImgUrl,
// 			GroupIcon:      group.GroupIcon,
// 			GroupIconColor: group.GroupIconColor,
// 			GroupType:      group.GroupType,
// 			GroupStatus:    group.GroupStatus,
// 			CreateAt:       group.CreateAt,
// 			UpdateAt:       group.UpdateAt,
// 		}
// 		groupResponses = append(groupResponses, groupResponse)
// 	}
// 	return groupResponses, nil
// }
