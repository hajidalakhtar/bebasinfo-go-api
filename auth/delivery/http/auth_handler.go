package http

import (
	"bebasinfo/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

func NewAuthHandler(app *fiber.App) {
	app.Get("/token/generate", Register)
}

func Register(c *fiber.Ctx) error {

	token, err := helper.GenerateToken()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{Message: err.Error()})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
