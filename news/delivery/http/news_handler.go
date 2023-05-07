package http

import (
	"bebasinfo/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
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
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	date := c.Query("date")
	page := c.Query("page", "1")
	pageInt, err := strconv.Atoi(page)

	limit := c.Query("limit", "10")
	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	news, paginate, err := b.NewsUsecase.Find(c.Context(), date, source, pageInt, limitInt)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.WebRespPaginate{
		Code:     http.StatusOK,
		Status:   "OK",
		Message:  "Success get news",
		Data:     news,
		Paginate: paginate,
	})

}

func (b NewsHandler) StoreFromRSS(c *fiber.Ctx) error {
	source := c.Query("source")
	if source == "" {
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	news, err := b.NewsUsecase.Store(c.Context(), source)
	if err != nil {
		fmt.Print(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success store news from RSS",
		Data:    news,
	})
}
