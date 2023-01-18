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
