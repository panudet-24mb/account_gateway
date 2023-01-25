package handler

import (
	"account_services/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) CreateNewUserAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.NewUserAccountRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	result, err := h.userService.NewUserAccount(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result.UserName,
	})

}

func (h userHandler) GetAccounts(c *fiber.Ctx) error {
	result, err := h.userService.GetAccounts()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userHandler) GetAccountByID(c *fiber.Ctx) error {
	id := c.Params("_id")
	result, err := h.userService.GetAccountByID(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (h userHandler) UpdateAccounts(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.UpdateUserRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userService.UpdateAccounts(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (h userHandler) UpdateAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")
	id := c.Params("_id")
	body := new(services.UpdateUserRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userService.UpdateAccount(id, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (h userHandler) DeleteAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")
	userName := c.Params("username")
	body := new(services.DeleteAccountRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userService.DeleteAccount(userName, body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Deleted Success",
		"data":   result.UserName,
	})
}

func (h userHandler) FindAccount(c *fiber.Ctx) error {
	body := new(services.FindAccountRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := h.userService.FindAccount(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}
