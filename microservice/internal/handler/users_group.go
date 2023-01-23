package handler

import (
	"account_gateway/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userGroupHandler struct {
	userGroupService services.UserGroupService
}

func NewUserGroupHandler(userGroupService services.UserGroupService) userGroupHandler {
	return userGroupHandler{userGroupService: userGroupService}
}

func (u userGroupHandler) CreateNewGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.NewGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userGroupService.NewGroup(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}

func (u userGroupHandler) GetGroups(c *fiber.Ctx) error {
	result, err := u.userGroupService.GetGroups()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (u userGroupHandler) GetGroupByID(c *fiber.Ctx) error {
	id := c.Params("_id")
	result, err := u.userGroupService.GetGroupByID(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (u userGroupHandler) UpdateGroups(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.UpdateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userGroupService.UpdateGroups(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (u userGroupHandler) UpdateGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("_id")
	body := new(services.UpdateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userGroupService.UpdateGroup(id, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (u userGroupHandler) DeleteGroup(c *fiber.Ctx) error {
	c.Accepts("application/json")
	groupName := c.Params("groupname")
	body := new(services.UpdateGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userGroupService.DeleteGroup(groupName, body)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Deleted Success",
		"data":   result.GroupName,
	})
}

func (u userGroupHandler) FindGroup(c *fiber.Ctx) error {
	body := new(services.FindGroupRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userGroupService.FindGroup(*body)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}
