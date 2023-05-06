package middleware

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type (
	ResponseError struct {
		Code        domain.ErrorCode `json:"code"`
		Description string           `json:"description"`
	}
)

func JWTMiddleware(c *fiber.Ctx) error {
	secretKey := viper.GetString(`secretKey`)

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Code:        domain.ErrUnauthorized,
			Description: "Unauthorized",
		})
	}
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Code:        domain.ErrUnauthorized,
			Description: "Unauthorized",
		})

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Code:        domain.ErrUnauthorized,
			Description: "Invalid or expired token",
		})
	}

	iss := claims["iss"]
	role := claims["role"]
	c.Context().SetUserValue("iss", iss)
	c.Context().SetUserValue("role", role)

	return c.Next()
}
