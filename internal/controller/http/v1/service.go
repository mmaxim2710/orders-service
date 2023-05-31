package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/controller/http/v1/middleware"
	"github.com/mmaxim2710/orders-service/internal/usecase"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"github.com/mmaxim2710/orders-service/pkg/validations"
)

type serviceRoutes struct {
	s usecase.Service
	l logger.Interface
}

func newServiceRoutes(handler fiber.Router, s usecase.Service, l logger.Interface) {
	r := &serviceRoutes{
		s: s,
		l: l,
	}

	h := handler.Group("/service")
	{
		h.Post("/create", middleware.Protected(), r.createService)
		h.Post("/update", middleware.Protected(), r.updateService)
		h.Get("/:serviceID", middleware.Protected(), r.serviceByID)
	}
}

type (
	doCreateServiceRequest struct {
		Title       string  `json:"title" validate:"required,min=3,max=128"`
		Description string  `json:"description" validate:"required,min=1,max=1024"`
		Price       float64 `json:"price" validate:"required,min=0"`
	}

	createServiceResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
	}
)

func (r *serviceRoutes) createService(ctx *fiber.Ctx) error {
	request := doCreateServiceRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.l.Error(err, "http - v1 - createService")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(request)
	if !ok {
		r.l.Error(ErrValidationFailed, "http - v1 - createService")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	jwtData := ctx.Locals("jwt").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	userID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "http - v1 - update")
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	service, err := r.s.Create(userID, request.Title, request.Description, request.Price)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := createServiceResponse{
		ID:          service.ID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
	}

	return successResponse(ctx, fiber.StatusOK, "Success create", response)
}

type (
	doUpdateServiceRequest struct {
		ID          string  `json:"id" validate:"required"`
		Title       string  `json:"title" validate:"required,min=3,max=128"`
		Description string  `json:"description" validate:"required,min=1,max=1024"`
		Price       float64 `json:"price" validate:"required,min=0"`
		IsClosed    bool    `json:"is_closed" validate:"required"`
	}

	updateServiceResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		IsClosed    bool      `json:"is_closed"`
	}
)

func (r *serviceRoutes) updateService(ctx *fiber.Ctx) error {
	request := doUpdateServiceRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.l.Error("Path: http - v1 - updateService. Error: %w", err)
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(&request)
	if !ok {
		r.l.Error("Path: http - v1 - updateService. Error: %w", ErrValidationFailed)
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	serviceID, err := uuid.Parse(request.ID)
	if err != nil {
		r.l.Error(err, "http - v1 - updateService. Error: %w", err)
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid format of uuid", err)
	}

	newUser, err := r.s.Update(serviceID, request.Title, request.Description, request.Price, request.IsClosed)
	if err != nil {
		r.l.Error("Path: http - v1 - updateService. Error: %w", err)
		if err == usecase.ErrServiceNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "Service with provided id not exists", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := &updateServiceResponse{
		ID:          newUser.ID,
		Title:       newUser.Title,
		Description: newUser.Description,
		Price:       newUser.Price,
		IsClosed:    newUser.IsClosed,
	}
	return successResponse(ctx, fiber.StatusOK, "Successful update", response)
}

type (
	serviceByIdResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		IsClosed    bool      `json:"is_closed"`
	}
)

func (r *serviceRoutes) serviceByID(ctx *fiber.Ctx) error {
	serviceID := ctx.Params("serviceID")

	serviceUUID, err := uuid.Parse(serviceID)
	if err != nil {
		r.l.Error(err, "http - v1 - serviceByID")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid format of uuid", err)
	}

	service, err := r.s.GetByID(serviceUUID)
	if err != nil {
		r.l.Error(err, "http - v1 - serviceByID")
		if err == usecase.ErrServiceNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "Service with provided id not exists", err.Error())
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := &serviceByIdResponse{
		ID:          service.ID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
		IsClosed:    service.IsClosed,
	}
	return successResponse(ctx, fiber.StatusOK, "Successful get", response)
}
