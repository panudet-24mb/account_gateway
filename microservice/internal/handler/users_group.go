package handler

import (
	"account_services/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userGroupHandler struct {
	userGroupService services.UserGroupService
}

func NewUserGroupHandler(userGroupService services.UserGroupService) userGroupHandler {
	return userGroupHandler{userGroupService: userGroupService}
}

func (h userGroupHandler) CreateNewGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.NewGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userGroupService.NewGroup(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}

func (h userGroupHandler) GetGroups(c *fiber.Ctx) error {
	result, err := h.userGroupService.GetGroups()
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userGroupHandler) GetGroupByID(c *fiber.Ctx) error {
	id := c.Params("_id")
	result, err := h.userGroupService.GetGroupByID(id)
	if err != nil {
		return fiber.ErrBadRequest
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userGroupHandler) UpdateGroups(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.UpdateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userGroupService.UpdateGroups(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (h userGroupHandler) UpdateGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("_id")
	body := new(services.UpdateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userGroupService.UpdateGroup(id, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (h userGroupHandler) DeleteGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	groupName := c.Params("groupname")
	body := new(services.UpdateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userGroupService.DeleteGroup(groupName, body)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Deleted Success",
		"data":   result.GroupName,
	})
}

// func (u userGroupHandler) FindGroup(c *fiber.Ctx) error {
// 	body := new(services.FindGroupRequest)
// 	if err := c.BodyParser(body); err != nil {
// 		return fiber.ErrBadRequest
// 	}
// 	result, err := u.userGroupService.FindGroup(*body)
// 	if err != nil {
// 		return err
// 	}
// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"data": result,
// 	})
// }
