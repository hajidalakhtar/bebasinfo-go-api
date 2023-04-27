package http

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/business/delivery/http/middleware"
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type (
	ResponseError struct {
		Code        domain.ErrorCode `json:"code"`
		Description string           `json:"description"`
	}
)

type BusinessHandler struct {
	BusinessUsecase domain.BusinessUsecase
}

func NewBusinessHandler(app *fiber.App, us domain.BusinessUsecase) {
	handler := &BusinessHandler{
		BusinessUsecase: us,
	}

	app.Post("/business", middleware.JWTMiddleware, middleware.AdminACL, handler.Store)
	app.Delete("/business/:id", middleware.JWTMiddleware, middleware.AdminACL, handler.Delete)
	app.Put("/business/:id", middleware.JWTMiddleware, middleware.AdminACL, handler.Update)
	app.Get("/business/search", handler.FindBusiness)

}

func (b BusinessHandler) Store(c *fiber.Ctx) error {
	var business domain.Business
	err := c.BodyParser(&business)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:        domain.ErrInvalidInput,
			Description: err.Error(),
		})
	}

	ctx := context.Background()
	err = b.BusinessUsecase.Store(ctx, &business)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:        domain.ErrInternal,
			Description: err.Error(),
		})
	}

	return c.JSON(business)
}

func (b BusinessHandler) Delete(c *fiber.Ctx) error {
	BusinessID := c.Params("id")
	BusinessUUID, err := uuid.Parse(BusinessID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:        domain.ErrInvalidInput,
			Description: err.Error(),
		})
	}

	ctx := context.Background()
	err = b.BusinessUsecase.Delete(ctx, BusinessUUID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:        domain.ErrInternal,
			Description: err.Error(),
		})
	}

	return c.JSON("Data berhasil dihapus")
}

func (b BusinessHandler) Update(c *fiber.Ctx) error {
	var business domain.Business

	businessID := c.Params("id")
	businessUUID, err := uuid.Parse(businessID)
	err = c.BodyParser(&business)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseError{
			Code:        domain.ErrInvalidInput,
			Description: err.Error(),
		})
	}

	ctx := context.Background()
	err = b.BusinessUsecase.Update(ctx, &business, businessUUID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:        domain.ErrInternal,
			Description: err.Error(),
		})
	}

	return c.JSON("Data berhasil diupdate")

}

func (b BusinessHandler) FindBusiness(c *fiber.Ctx) error {
	var businesses []domain.Business

	term := c.Query("term")
	sortBy := c.Query("sort_by")

	latitude := c.Query("latitude")
	latitudeFloat, ok := strconv.ParseFloat(latitude, 64)
	if ok != nil {
		latitudeFloat = 0
	}

	longitude := c.Query("longitude")
	longitudeFloat, ok := strconv.ParseFloat(longitude, 64)
	if ok != nil {
		longitudeFloat = 0
	}

	limit := c.Query("limit")
	limitInt, ok := strconv.Atoi(limit)
	if ok != nil {
		limitInt = 0
	}
	offset := c.Query("offset")
	offsetInt, ok := strconv.Atoi(offset)
	if ok != nil {
		offsetInt = 0
	}

	businesses, err := b.BusinessUsecase.Find(c.Context(), term, sortBy, limitInt, offsetInt, latitudeFloat, longitudeFloat)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ResponseError{
			Code:        domain.ErrInternal,
			Description: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"businesses": businesses,
	})
}
