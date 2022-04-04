package user

import (
	"github.com/busranurguner/foodchain/pkg/models"
	validation "github.com/busranurguner/foodchain/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetAll(c *fiber.Ctx) error
}

type userHandler struct {
	service UserService
}

var _ UserHandler = userHandler{}

func NewHandler(service UserService) UserHandler {
	return userHandler{service: service}
}

func (u userHandler) GetAll(c *fiber.Ctx) error {
	request := GetAllRequest{}
	err := c.QueryParser(&request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = validation.Validator.Struct(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	user, err := u.service.GetAll(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	return c.Status(200).JSON(models.Response{Data: user})

}
