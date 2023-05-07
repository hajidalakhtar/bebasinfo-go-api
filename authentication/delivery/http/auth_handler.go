package http

import (
	"bebasinfo/domain"
	"bebasinfo/user/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(app *fiber.App, us domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: us,
	}
	app.Post("/auth/register", handler.Register)
	app.Post("/auth/login", handler.Login)
	app.Get("/auth/me", middleware.JWTMiddleware, handler.GetMe)
}

func (u *AuthHandler) Register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	username := c.FormValue("username")

	if email == "" || password == "" || username == "" {
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "email, password, and username is required",
		})
	}

	result, err := u.AuthUsecase.Register(c.Context(), username, password, email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: err.Error(),
		})
	}
	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Message: "Success register user",
		Data:    result,
	})
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	result, isSuccess, _ := u.AuthUsecase.Login(c.Context(), email, password)
	if !isSuccess {
		return c.Status(http.StatusUnauthorized).JSON(domain.WebResponse{
			Code:    http.StatusUnauthorized,
			Status:  domain.ErrUnauthorized,
			Message: "Email or Password is wrong",
		})
	}
	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Message: "Success login user",
		Data:    result,
	})
}

func (u *AuthHandler) GetMe(c *fiber.Ctx) error {
	userID, ok := c.Context().Value("userID").(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: "Error when parsing userID",
		})
	}

	userUUID, _ := uuid.Parse(userID)

	result, err := u.AuthUsecase.GetMe(c.Context(), userUUID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: err.Error(),
		})

	}

	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get user",
		Data:    result,
	})
}
