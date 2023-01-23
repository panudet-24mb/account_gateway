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

func (u userHandler) GetAccounts(c *fiber.Ctx) error {
	result, err := u.userService.GetAccounts()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (u userHandler) GetAccountByID(c *fiber.Ctx) error {
	id := c.Params("_id")
	result, err := u.userService.GetAccountByID(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}

func (u userHandler) UpdateAccounts(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(services.UpdateUserRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userService.UpdateAccounts(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   result,
	})
}

func (u userHandler) UpdateAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")
	username := c.Params("username")
	body := new(services.UpdateUserRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	_, err := u.userService.UpdateAccount(body)
	if err != nil {
		return handleError(c, err)
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Updated Success",
		"data":   username,
	})
}

func (u userHandler) DeleteAccount(c *fiber.Ctx) error {
	c.Accepts("application/json")
	userName := c.Params("username")
	body := new(services.DeleteAccountRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userService.DeleteAccount(userName, body)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "Deleted Success",
		"data":   result.UserName,
	})
}

func (u userHandler) FindAccount(c *fiber.Ctx) error {
	body := new(services.FindAccountRequest)
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}
	result, err := u.userService.FindAccount(*body)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": result,
	})
}
