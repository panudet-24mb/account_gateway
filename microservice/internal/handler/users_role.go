package handler

import (
	"account_services/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type roleHandler struct {
	roleService services.RoleService
}

func NewRoleHandler(roleService services.RoleService) roleHandler {
	return roleHandler{roleService: roleService}
}

func (h roleHandler) CreateNewRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.NewRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.roleService.NewRole(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h roleHandler) GetRoles(c *fiber.Ctx) error {
	result, err := h.roleService.GetRoles()
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h roleHandler) GetRoleByID(c *fiber.Ctx) error {
	id := c.Params("_id")
	result, err := h.roleService.GetRoleByID(id)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h roleHandler) UpdateRoles(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.UpdateRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.roleService.UpdateRoles(body)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})

}

func (h roleHandler) UpdateRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("_id")
	body := new(services.UpdateRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.roleService.UpdateRole(id, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})

}

func (h roleHandler) DeleteRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	userName := c.Params("username")
	body := new(services.UpdateRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.roleService.DeleteRole(userName, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Deleted Success",
		"data":   result.RoleName,
	})
}

type userRoleHandler struct {
	userRoleService services.UserRoleService
}

func NewUserRoleHandler(userRoleService services.UserRoleService) userRoleHandler {
	return userRoleHandler{userRoleService: userRoleService}
}

func (h userRoleHandler) CreateNewUserRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.NewUserRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userRoleService.NewUserRole(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userRoleHandler) GetUserRole(c *fiber.Ctx) error {
	result, err := h.userRoleService.GetUsersRole()
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userRoleHandler) GetUserRoleByID(c *fiber.Ctx) error {
	id := c.Params("_id")
	result, err := h.userRoleService.GetUserRoleByID(id)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userRoleHandler) UpdateUsersRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.UpdateUserRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userRoleService.UpdateUserRoles(body)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})

}

func (h userRoleHandler) UpdateUserRole(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("_id")
	body := new(services.UpdateUserRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userRoleService.UpdateUserRole(id, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})

}

func (h userRoleHandler) DeleteUserRole(c *fiber.Ctx) error {
	id := c.Params("_id")
	body := new(services.DeleteUserRoleRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userRoleService.DeleteUserRole(id, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Deleted Success",
		"data":   result.ObjID,
	})
}

// func (h userRoleHandler) DeleteUserRole(c *fiber.Ctx) error {
// 	id := c.Params("_id")
// 	result, err := h.userRoleService.DeleteUserRole(id)
// 	if err != nil {
// 		return handleError(c, err)
// 	}
// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"status": "Deleted Success",
// 		"data":   result,
// 	})
// }
