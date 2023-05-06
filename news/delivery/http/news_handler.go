package http

import (
	"bebasinfo/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	ResponseError struct {
		Code    domain.ErrorCode `json:"code"`
		Message string           `json:"message"`
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

}

func (b NewsHandler) FindNews(c *fiber.Ctx) error {
	var news []domain.News

	source := c.Query("source")
	if source == "" {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:    domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	date := c.Query("date")
	news, err := b.NewsUsecase.Find(c.Context(), date, source)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:    domain.ErrInternal,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Message: "Success get news",
		Data:    news,
	})

}

func (b NewsHandler) StoreFromRSS(c *fiber.Ctx) error {
	source := c.Query("source")
	if source == "" {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:    domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	news, err := b.NewsUsecase.Store(c.Context(), source)
	if err != nil {
		fmt.Print(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:    domain.ErrInternal,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Message: "Success store news from RSS",
		Data:    news,
	})
}
