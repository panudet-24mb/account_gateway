package services

type NewRoleRequest struct {
	RoleName      string `json:"rolename"`
	RoleCode      string `json:"rolecode"`
	RoleDesc      string `json:"roledesc"`
	RoleImgUrl    string `json:"roleimgurl"`
	RoleIcon      string `json:"roleicon"`
	RoleIconColor string `json:"roleiconcolor"`
	RoleType      string `json:"roletype"`
	RoleStatus    string `json:"rolestatus"`
	RoleActive    bool   `json:"roleactive"`
}

type NewRoleResponse struct {
	RoleName      string `json:"rolename"`
	RoleCode      string `json:"rolecode"`
	RoleDesc      string `json:"roledesc"`
	RoleImgUrl    string `json:"roleimgurl"`
	RoleIcon      string `json:"roleicon"`
	RoleIconColor string `json:"roleiconcolor"`
	RoleType      string `json:"roletype"`
	RoleStatus    string `json:"rolestatus"`
	RoleActive    bool   `json:"roleactive"`
	CreateAt      string `json:"createat"`
	UpdateAt      string `json:"updateat"`
	DeleteAt      string `json:"deleteat"`
}

type UpdateRoleRequest struct {
	RoleName      string `json:"rolename"`
	RoleCode      string `json:"rolecode"`
	RoleDesc      string `json:"roledesc"`
	RoleImgUrl    string `json:"roleimgurl"`
	RoleIcon      string `json:"roleicon"`
	RoleIconColor string `json:"roleiconcolor"`
	RoleType      string `json:"roletype"`
	RoleStatus    string `json:"rolestatus"`
	RoleActive    bool   `json:"roleactive"`
}

type UpdateRoleResponse struct {
	RoleName      string `json:"rolename"`
	RoleCode      string `json:"rolecode"`
	RoleDesc      string `json:"roledesc"`
	RoleImgUrl    string `json:"roleimgurl"`
	RoleIcon      string `json:"roleicon"`
	RoleIconColor string `json:"roleiconcolor"`
	RoleType      string `json:"roletype"`
	RoleStatus    string `json:"rolestatus"`
	RoleActive    bool   `json:"roleactive"`
	CreateAt      string `json:"createat"`
	UpdateAt      string `json:"updateat"`
	DeleteAt      string `json:"deleteat"`
}

type FindRoleRequest struct {
	RoleName      string `json:"rolename"`
	RoleCode      string `json:"rolecode"`
	RoleDesc      string `json:"roledesc"`
	RoleImgUrl    string `json:"roleimgurl"`
	RoleIcon      string `json:"roleicon"`
	RoleIconColor string `json:"roleiconcolor"`
	RoleType      string `json:"roletype"`
	RoleStatus    string `json:"rolestatus"`
	RoleActive    bool   `json:"roleactive"`
}

type FindRoleResponse struct {
	RoleName      string `json:"rolename"`
	RoleCode      string `json:"rolecode"`
	RoleDesc      string `json:"roledesc"`
	RoleImgUrl    string `json:"roleimgurl"`
	RoleIcon      string `json:"roleicon"`
	RoleIconColor string `json:"roleiconcolor"`
	RoleType      string `json:"roletype"`
	RoleStatus    string `json:"rolestatus"`
	RoleActive    bool   `json:"roleactive"`
	CreateAt      string `json:"createat"`
	UpdateAt      string `json:"updateat"`
	DeleteAt      string `json:"deleteat"`
}

type RoleService interface {
	NewRole(newRoleRequest *NewRoleRequest) (*NewRoleResponse, error)
	GetRoles() ([]NewRoleResponse, error)
	GetRoleByID(id string) (*NewRoleResponse, error)
	UpdateRoles(updateRoleRequest *UpdateRoleRequest) (*UpdateRoleResponse, error)
	UpdateRole(id string, updateRoleRequest *UpdateRoleRequest) (*UpdateRoleResponse, error)
	DeleteRole(roleName string, updateRoleRequest *UpdateRoleRequest) (*UpdateRoleResponse, error)
	// FindRole(findRole FindRoleRequest) ([]FindRoleResponse, error)

}

type NewUserRoleRequest struct {
	UserID string `json:"userid"`
	RoleID string `json:"roleid"`
}

type NewUserRoleResponse struct {
	UserID   string `json:"userid"`
	RoleID   string `json:"roleid"`
	CreateAt string `json:"createat"`
	UpdateAt string `json:"updateat"`
}

type UpdateUserRoleRequest struct {
	UserID   string `json:"userid"`
	RoleID   string `json:"roleid"`
	CreateAt string `json:"createat"`
	UpdateAt string `json:"updateat"`
	DeleteAt string `json:"deleteat"`
}

type UpdateUserRoleResponse struct {
	UserID   string `json:"userid"`
	RoleID   string `json:"roleid"`
	CreateAt string `json:"createat"`
	UpdateAt string `json:"updateat"`
	DeleteAt string `json:"deleteat"`
}

type DeleteUserRoleRequest struct {
	ObjID    string `json:"_id"`
	DeleteAt string `json:"deleteat"`
}

type DeleteUserRoleResponse struct {
	ObjID    string `json:"_id"`
	DeleteAt string `json:"deleteat"`
}

type UserRoleService interface {
	NewUserRole(newUserRole *NewUserRoleRequest) (*NewUserRoleResponse, error)
	GetUsersRole() ([]NewUserRoleResponse, error)
	GetUserRoleByID(id string) (*NewUserRoleResponse, error)
	UpdateUserRoles(updateUsersRole *UpdateUserRoleRequest) (*UpdateUserRoleResponse, error)
	UpdateUserRole(id string, updateUserRole *UpdateUserRoleRequest) (*UpdateUserRoleResponse, error)
	DeleteUserRole(id string, deleteUserRole *DeleteUserRoleRequest) (*DeleteUserRoleResponse, error)
}
