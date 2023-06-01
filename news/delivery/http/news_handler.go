package http

import (
	"bebasinfo/domain"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

type NewsHandler struct {
	NewsUsecase domain.NewsUsecase
	Logger      *zap.Logger
}

func NewNewsHandler(app *fiber.App, ns domain.NewsUsecase, logger *zap.Logger) {
	handler := &NewsHandler{
		NewsUsecase: ns,
		Logger:      logger,
	}

	app.Get("/news/detail/:id", handler.Find)
	app.Get("/news/search", handler.FindNews)
	app.Get("/store/news", handler.Store)
	app.Get("/store/multiple/newsdata", handler.StoreMultipleNewsData)

}

func (b NewsHandler) FindNews(c *fiber.Ctx) error {
	var news []domain.News
	var categoryArr []string

	source := c.Query("source")
	category := c.Query("category")
	if source == "" {
		return c.Status(http.StatusBadRequest).JSON(domain.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  domain.ErrInvalidInput,
			Message: "source is required",
		})
	}

	sourceArr := strings.Split(source, ",")
	if category != "" {
		categoryArr = strings.Split(category, ",")
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

	news, paginate, err := b.NewsUsecase.Search(c.Context(), date, sourceArr, categoryArr, pageInt, limitInt)
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

	sourceArr := strings.Split(source, ",")

	news, err := b.NewsUsecase.Store(c.Context(), newsResource, category, sourceArr)
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
		Message: "Success store news from RSS",
		Data:    news,
	})
}

func (b NewsHandler) StoreMultipleNewsData(c *fiber.Ctx) error {

	category := c.Query("category")
	news, err := b.NewsUsecase.StoreMultiplePagesFromNewsDataApi(c.Context(), category)
	if err != nil {
		c.JSON(domain.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  domain.ErrInternal,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success store news from News Data API",
		Data:    news,
	})
}
