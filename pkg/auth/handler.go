package auth

import (
	"github.com/busranurguner/foodchain/pkg/models"
	validation "github.com/busranurguner/foodchain/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SignUp(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}

type authHandler struct {
	service AuthService
}

var _ AuthHandler = authHandler{}

func NewHandler(service AuthService) AuthHandler {
	return authHandler{service: service}
}

func (a authHandler) SignUp(c *fiber.Ctx) error {

	request := SignUpRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = validation.Validator.Struct(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = a.service.SignUp(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	return c.SendStatus(200)
}

func (a authHandler) Login(c *fiber.Ctx) error {
	request := LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = validation.Validator.Struct(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	at, rt, err := a.service.Login(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	response := map[string]string{
		"access":  at,
		"refresh": rt,
	}
	return c.Status(200).JSON(models.Response{Data: response})
}

func (a authHandler) Refresh(c *fiber.Ctx) error {
	request := RefreshRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = validation.Validator.Struct(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	at, rt, err := a.service.Refresh(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	response := map[string]string{
		"access":  at,
		"refresh": rt,
	}
	return c.Status(200).JSON(models.Response{Data: response})
}
