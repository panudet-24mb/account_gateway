package handler

import (
	"account_gateway/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (u userHandler) CreateNewUserAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.NewUserAccountRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	result, err := u.userService.NewUserAccount(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result.UserName,
	})

}
