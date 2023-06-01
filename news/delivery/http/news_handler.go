package http

import (
	"bebasinfo/domain"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type NewsHandler struct {
	NewsUsecase domain.NewsUsecase
	Logger      *zap.Logger
}

var (
	cronJob *cron.Cron
	mutex   sync.Mutex
)

func NewNewsHandler(app *fiber.App, ns domain.NewsUsecase, logger *zap.Logger) {
	handler := &NewsHandler{
		NewsUsecase: ns,
		Logger:      logger,
	}

	app.Get("/news/cronjob/:status", handler.CronJob)

	app.Get("/news/detail/:id", handler.Find)
	app.Get("/news/search", handler.FindNews)
	app.Get("/store/news", handler.Store)

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

	sourceArr := strings.Split(source, ",")
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

	news, paginate, err := b.NewsUsecase.Search(c.Context(), date, sourceArr, pageInt, limitInt)
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

func (b NewsHandler) CronJob(c *fiber.Ctx) error {
	status := c.Params("status")
	if status == "start" {
		startCronJob(b.Logger)
	}

	if status == "stop" {
		stopCronJob()
	}
	return c.JSON("cron job : " + status)
}

func hello() {
	fmt.Println("Hello, World!")
}

func startCronJob(log *zap.Logger) {

	log.Info("Create new cron")
	cronJob := cron.New(cron.WithSeconds())
	cronJob.AddFunc("*/1 * * * *", func() {
		log.Info("cron")
	})

	log.Info("Start cron")
	cronJob.Start()

	//printCronEntries(c.Entries())
	//time.Sleep(2 * time.Minute)
	//
	//// Funcs may also be added to a running Cron
	//log.Info("Add new job to a running cron")
	//entryID2, _ := c.AddFunc("*/2 * * * *", func() { log.Info("[Job 2]Every two minutes job\n") })
	//printCronEntries(c.Entries())
	//time.Sleep(5 * time.Minute)
	//
	////Remove Job2 and add new Job2 that run every 1 minute
	//log.Info("Remove Job2 and add new Job2 with schedule run every minute")
	//c.Remove(entryID2)
	//c.AddFunc("*/1 * * * *", func() { log.Info("[Job 2]Every one minute job\n") })
	//time.Sleep(5 * time.Minute)

}

func stopCronJob() {
	cronJob.Stop()

	//scheduler.Clear()
	//scheduler.Stop()
}
