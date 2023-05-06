package middleware

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"github.com/gofiber/fiber/v2"
)

func AdminACL(c *fiber.Ctx) error {
	role := c.Context().Value("role").(float64)
	if role != 1 {
		return c.Status(fiber.StatusForbidden).JSON(ResponseError{
			Code:        domain.ErrForbidden,
			Description: "forbidden",
		})
	}
	return c.Next()
}
