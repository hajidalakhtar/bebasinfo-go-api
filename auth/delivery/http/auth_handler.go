package http

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/helper"
	//"eventzezz_backend/user/delivery/http/middleware"
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

//func GetMe(c *fiber.Ctx) error {
//	//userID, ok := c.Context().Value("userID").(string)
//	//if !ok {
//	//	return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: "Error when parsing userID"})
//	//}
//	//
//	//userUUID, _ := uuid.Parse(userID)
//	//
//	//result, err := u.AuthUsecase.GetMe(c.Context(), userUUID)
//	//if err != nil {
//	//	return c.JSON(domain.WebResponse{Code: 500, Status: "Error", Data: err.Error()})
//	//}
//	//
//	//return c.JSON(domain.WebResponse{Code: http.StatusOK, Status: "OK", Data: result})
//}
