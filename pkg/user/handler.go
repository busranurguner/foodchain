package user

import (
	"github.com/busranurguner/foodchain/pkg/models"
	validation "github.com/busranurguner/foodchain/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
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
func (u userHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(models.Response{Error: "id path parameter can not be empty"})
	}
	user, err := u.service.GetByID(id)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	return c.Status(200).JSON(models.Response{Data: user})
}

//Update password
func (u userHandler) Update(c *fiber.Ctx) error {
	req := UpdateRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	req.ID = c.Params("id")
	err = validation.Validator.Struct(req)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = u.service.Update(req)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	return c.SendStatus(200)
}

func (u userHandler) Delete(c *fiber.Ctx) error {
	request := DeleteRequest{
		ID: c.Params("id"),
	}
	err := validation.Validator.Struct(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	err = u.service.Delete(request)
	if err != nil {
		return c.Status(400).JSON(models.Response{Error: err.Error()})
	}
	return c.SendStatus(200)
}
