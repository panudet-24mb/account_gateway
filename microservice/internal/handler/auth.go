package handler

import (
	"account_gateway/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authSrv services.AuthService
}

func NewAuthHandler(authSrv services.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) DefaultLogin(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.LoginDefaultForm)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.authSrv.DefaultLogin(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})

}