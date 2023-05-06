package http

import (
	"bebasinfo/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	ResponseError struct {
		Code        domain.ErrorCode `json:"code"`
		Description string           `json:"description"`
	}
)

type NewsHandler struct {
	NewsUsecase domain.NewsUsecase
}

func NewNewsHandler(app *fiber.App, ns domain.NewsUsecase) {
	handler := &NewsHandler{
		NewsUsecase: ns,
	}

	app.Get("/news/search", handler.FindNews)
	app.Get("/store/rss/news", handler.StoreFromRSS)

	//app.Post("/business", middleware.JWTMiddleware, middleware.AdminACL, handler.Store)
	//app.Delete("/business/:id", middleware.JWTMiddleware, middleware.AdminACL, handler.Delete)
	//app.Put("/business/:id", middleware.JWTMiddleware, middleware.AdminACL, handler.Update)

}

func (b NewsHandler) FindNews(c *fiber.Ctx) error {
	var news []domain.News

	source := c.Query("source")
	if source == "" {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:        domain.ErrInvalidInput,
			Description: "source is required",
		})
	}

	date := c.Query("date")
	news, err := b.NewsUsecase.Find(c.Context(), date, source)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:        domain.ErrInternal,
			Description: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"businesses": news,
	})
}

func (b NewsHandler) StoreFromRSS(c *fiber.Ctx) error {
	source := c.Query("source")
	if source == "" {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:        domain.ErrInvalidInput,
			Description: "source is required",
		})
	}

	news, err := b.NewsUsecase.Store(c.Context(), source)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:        domain.ErrInternal,
			Description: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": news,
	})
}
