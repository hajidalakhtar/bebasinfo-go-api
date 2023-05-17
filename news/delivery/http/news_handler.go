package http

import (
	"bebasinfo/domain"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	app.Get("/news/detail/:id", handler.Find)
	app.Get("/news/search", handler.FindNews)
	app.Get("/store/rss/news", handler.Store)

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

	news, paginate, err := b.NewsUsecase.Search(c.Context(), date, source, pageInt, limitInt)
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

func (b NewsHandler) Find(c *fiber.Ctx) error {
	var news []domain.News
	newsIdStr := c.Params("id")
	if newsIdStr == "" {
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "newsId is required",
		})
	}

	newsUUID, err := uuid.Parse(newsIdStr)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "newsId is required",
		})
	}

	news, err = b.NewsUsecase.Find(c.Context(), newsUUID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: err.Error(),
		})
	}

	if len(news) == 0 {
		return c.Status(http.StatusInternalServerError).JSON(domain.WebResponse{
			Code:    http.StatusNotFound,
			Status:  domain.ErrNotFound,
			Message: "News Tidak ditemukan",
		})
	}

	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get news",
		Data:    news[0],
	})

}

func (b NewsHandler) Store(c *fiber.Ctx) error {
	newsResource := c.Query("news_resource")
	category := c.Query("category")
	source := c.Query("source")

	validate := validator.New()

	// Memeriksa apakah newsResource tidak kosong
	if err := validate.Var(newsResource, "required"); err != nil {
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	// Memeriksa apakah category tidak kosong jika newsResource bernilai "api"
	if newsResource == "api" {
		if err := validate.Var(category, "required"); err != nil {
			return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  domain.ErrInvalidInput,
				Message: "source is required",
			})
		}
	}

	// Memeriksa apakah source tidak kosong jika newsResource bernilai "rss"
	if newsResource == "rss" {
		if err := validate.Var(source, "required"); err != nil {
			return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  domain.ErrInvalidInput,
				Message: "source is required",
			})
		}
	}

	news, err := b.NewsUsecase.Store(c.Context(), newsResource, category, source)
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
