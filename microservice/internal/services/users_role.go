package services

import (
	"account_services/internal/repository"
	"strconv"
	"time"
)

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepository repository.RoleRepository) RoleService {
	return roleService{roleRepository: roleRepository}
}

func (s roleService) NewRole(newRoleRequest *NewRoleRequest) (*NewRoleResponse, error) {
	role := repository.Roles{
		RoleName:      newRoleRequest.RoleName,
		RoleCode:      newRoleRequest.RoleCode,
		RoleDesc:      newRoleRequest.RoleDesc,
		RoleImgUrl:    newRoleRequest.RoleImgUrl,
		RoleIcon:      newRoleRequest.RoleIcon,
		RoleIconColor: newRoleRequest.RoleIconColor,
		RoleType:      newRoleRequest.RoleType,
		RoleStatus:    newRoleRequest.RoleStatus,
		RoleActive:    newRoleRequest.RoleActive,
		CreateAt:      strconv.Itoa(int(time.Now().Unix())),
	}
	data, err := s.roleRepository.Create(role)
	if err != nil {
		return nil, err
	}

	roleResponse := &NewRoleResponse{
		RoleName:      data.RoleName,
		RoleCode:      data.RoleCode,
		RoleDesc:      data.RoleDesc,
		RoleImgUrl:    data.RoleImgUrl,
		RoleIcon:      data.RoleIcon,
		RoleIconColor: data.RoleIconColor,
		RoleType:      data.RoleType,
		RoleStatus:    data.RoleStatus,
		RoleActive:    data.RoleActive,
		CreateAt:      data.CreateAt,
	}
	return roleResponse, nil
}

func (s roleService) GetRoles() ([]NewRoleResponse, error) {
	roles, err := s.roleRepository.GetAll()
	if err != nil {
		return nil, err
	}
	roleResponses := []NewRoleResponse{}
	for _, role := range roles {
		roleResponse := NewRoleResponse{
			RoleName:      role.RoleName,
			RoleCode:      role.RoleCode,
			RoleDesc:      role.RoleDesc,
			RoleImgUrl:    role.RoleImgUrl,
			RoleIcon:      role.RoleIcon,
			RoleIconColor: role.RoleIconColor,
			RoleType:      role.RoleType,
			RoleStatus:    role.RoleStatus,
			RoleActive:    role.RoleActive,
			CreateAt:      role.CreateAt,
			UpdateAt:      role.UpdateAt,
		}
		roleResponses = append(roleResponses, roleResponse)
	}
	return roleResponses, nil
}

func (s roleService) GetRoleByID(id string) (*NewRoleResponse, error) {
	role, err := s.roleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	roleResponse := NewRoleResponse{
		RoleName:      role.RoleName,
		RoleCode:      role.RoleCode,
		RoleDesc:      role.RoleDesc,
		RoleImgUrl:    role.RoleImgUrl,
		RoleIcon:      role.RoleIcon,
		RoleIconColor: role.RoleIconColor,
		RoleType:      role.RoleType,
		RoleStatus:    role.RoleStatus,
		RoleActive:    role.RoleActive,
		CreateAt:      role.CreateAt,
		UpdateAt:      role.UpdateAt,
	}
	return &roleResponse, nil
}

func (s roleService) UpdateRoles(updateRoleRequest *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	roleData := repository.Roles{
		RoleName:      updateRoleRequest.RoleName,
		RoleCode:      updateRoleRequest.RoleCode,
		RoleDesc:      updateRoleRequest.RoleDesc,
		RoleImgUrl:    updateRoleRequest.RoleImgUrl,
		RoleIcon:      updateRoleRequest.RoleIcon,
		RoleIconColor: updateRoleRequest.RoleIconColor,
		RoleType:      updateRoleRequest.RoleType,
		RoleStatus:    updateRoleRequest.RoleStatus,
		RoleActive:    updateRoleRequest.RoleActive,
		UpdateAt:      strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.roleRepository.Update(roleData)
	if err != nil {
		return nil, err
	}
	roleUpdate := &UpdateRoleResponse{
		RoleName:      dataUpdate.RoleName,
		RoleCode:      dataUpdate.RoleCode,
		RoleDesc:      dataUpdate.RoleDesc,
		RoleImgUrl:    dataUpdate.RoleImgUrl,
		RoleIcon:      dataUpdate.RoleIcon,
		RoleIconColor: dataUpdate.RoleIconColor,
		RoleType:      dataUpdate.RoleType,
		RoleStatus:    dataUpdate.RoleStatus,
		RoleActive:    dataUpdate.RoleActive,
	}
	return roleUpdate, nil
}

func (s roleService) UpdateRole(id string, updateRoleRequest *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	roleData := repository.Roles{
		RoleName:      updateRoleRequest.RoleName,
		RoleCode:      updateRoleRequest.RoleCode,
		RoleDesc:      updateRoleRequest.RoleDesc,
		RoleImgUrl:    updateRoleRequest.RoleImgUrl,
		RoleIcon:      updateRoleRequest.RoleIcon,
		RoleIconColor: updateRoleRequest.RoleIconColor,
		RoleType:      updateRoleRequest.RoleType,
		RoleStatus:    updateRoleRequest.RoleStatus,
		RoleActive:    updateRoleRequest.RoleActive,
		UpdateAt:      strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.roleRepository.UpdateOne(id, roleData)
	if err != nil {
		return nil, err
	}
	roleUpdate := &UpdateRoleResponse{
		RoleName:      dataUpdate.RoleName,
		RoleCode:      dataUpdate.RoleCode,
		RoleDesc:      dataUpdate.RoleDesc,
		RoleImgUrl:    dataUpdate.RoleImgUrl,
		RoleIcon:      dataUpdate.RoleIcon,
		RoleIconColor: dataUpdate.RoleIconColor,
		RoleType:      dataUpdate.RoleType,
		RoleStatus:    dataUpdate.RoleStatus,
		RoleActive:    dataUpdate.RoleActive,
	}
	return roleUpdate, nil
}

func (s roleService) DeleteRole(roleName string, updateRoleRequest *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	roleData := repository.Roles{
		RoleName: updateRoleRequest.RoleName,
		DeleteAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataDelete, err := s.roleRepository.DeleteOne(roleName, roleData)
	if err != nil {
		return nil, err
	}
	roleDelelte := &UpdateRoleResponse{
		RoleName: dataDelete.RoleName,
		DeleteAt: dataDelete.DeleteAt,
	}
	return roleDelelte, nil
}

type userRoleService struct {
	userRoleRepository repository.UserRoleRepository
}

func NewUserRoleService(userRoleRepository repository.UserRoleRepository) UserRoleService {
	return userRoleService{userRoleRepository: userRoleRepository}
}

func (s userRoleService) NewUserRole(newUserRole *NewUserRoleRequest) (*NewUserRoleResponse, error) {
	userRole := repository.UsersRole{
		UserID:   newUserRole.UserID,
		RoleID:   newUserRole.RoleID,
		CreateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	data, err := s.userRoleRepository.Create(userRole)
	if err != nil {
		return nil, err
	}

	userRoleResponse := &NewUserRoleResponse{
		UserID:   data.UserID,
		RoleID:   data.RoleID,
		CreateAt: data.CreateAt,
	}

	return userRoleResponse, nil

}

func (s userRoleService) GetUsersRole() ([]NewUserRoleResponse, error) {
	usersRole, err := s.userRoleRepository.GetAll()
	if err != nil {
		return nil, err
	}
	usersRoleResponses := []NewUserRoleResponse{}
	for _, userRole := range usersRole {
		usersRoleResponse := NewUserRoleResponse{
			UserID:   userRole.UserID,
			RoleID:   userRole.RoleID,
			CreateAt: userRole.CreateAt,
			UpdateAt: userRole.UpdateAt,
		}
		usersRoleResponses = append(usersRoleResponses, usersRoleResponse)
	}
	return usersRoleResponses, nil
}

func (s userRoleService) GetUserRoleByID(id string) (*NewUserRoleResponse, error) {
	userRole, err := s.userRoleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	userRoleResponse := NewUserRoleResponse{
		UserID:   userRole.UserID,
		RoleID:   userRole.RoleID,
		CreateAt: userRole.CreateAt,
		UpdateAt: userRole.UpdateAt,
	}
	return &userRoleResponse, nil
}

func (s userRoleService) UpdateUserRoles(updateUsersRole *UpdateUserRoleRequest) (*UpdateUserRoleResponse, error) {
	userRoleData := repository.UsersRole{
		UserID:   updateUsersRole.UserID,
		RoleID:   updateUsersRole.RoleID,
		UpdateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.userRoleRepository.Update(userRoleData)
	if err != nil {
		return nil, err
	}
	userRoleUpdate := &UpdateUserRoleResponse{
		UserID:   dataUpdate.UserID,
		RoleID:   dataUpdate.RoleID,
		UpdateAt: dataUpdate.UpdateAt,
	}
	return userRoleUpdate, nil
}

func (s userRoleService) UpdateUserRole(id string, updateUserRole *UpdateUserRoleRequest) (*UpdateUserRoleResponse, error) {
	userRoleData := repository.UsersRole{
		UserID:   updateUserRole.UserID,
		RoleID:   updateUserRole.RoleID,
		UpdateAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataUpdate, err := s.userRoleRepository.UpdateOne(id, userRoleData)
	if err != nil {
		return nil, err
	}
	userRoleUpdate := &UpdateUserRoleResponse{
		UserID:   dataUpdate.UserID,
		RoleID:   dataUpdate.RoleID,
		UpdateAt: dataUpdate.UpdateAt,
	}
	return userRoleUpdate, nil
}

func (s userRoleService) DeleteUserRole(id string, deleteUserRole *DeleteUserRoleRequest) (*DeleteUserRoleResponse, error) {
	userRoleData := repository.UsersRole{
		DeleteAt: strconv.Itoa(int(time.Now().Unix())),
	}
	dataDelete, err := s.userRoleRepository.DeleteOne(id, userRoleData)
	if err != nil {
		return nil, err
	}
	userRoleDelete := &DeleteUserRoleResponse{
		ObjID:    id,
		DeleteAt: dataDelete.DeleteAt,
	}
	return userRoleDelete, nil
}
