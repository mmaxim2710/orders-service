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

	h := handler.Group("/services", middleware.Protected())
	{
		h.Post("", r.createService)
		h.Patch("", r.updateService)
		h.Get("/:serviceID", r.serviceByID)
		h.Get("", r.servicesByUserID)
		h.Delete("", r.deleteService)
	}
}

type (
	doCreateServiceRequest struct {
		Title       string  `json:"title" validate:"required,min=3,max=128"`
		Description string  `json:"description" validate:"required,min=1,max=1024"`
		Price       float64 `json:"price" validate:"required,min=0"`
	}

	doUpdateServiceRequest struct {
		ID          string  `json:"id" validate:"required,uuid4"`
		Title       string  `json:"title" validate:"required,min=3,max=128"`
		Description string  `json:"description" validate:"required,min=1,max=1024"`
		Price       float64 `json:"price" validate:"required,min=0"`
		IsClosed    bool    `json:"is_closed" validate:"required"`
	}

	doDeleteServiceRequest struct {
		ID string `json:"id" validate:"required,uuid4"`
	}

	serviceResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		IsClosed    bool      `json:"is_closed"`
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

	response := &serviceResponse{
		ID:          service.ID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
		IsClosed:    service.IsClosed,
	}

	return successResponse(ctx, fiber.StatusOK, "Success create", response)
}

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

	response := &serviceResponse{
		ID:          newUser.ID,
		Title:       newUser.Title,
		Description: newUser.Description,
		Price:       newUser.Price,
		IsClosed:    newUser.IsClosed,
	}
	return successResponse(ctx, fiber.StatusOK, "Successful update", response)
}

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

	response := &serviceResponse{
		ID:          service.ID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
		IsClosed:    service.IsClosed,
	}
	return successResponse(ctx, fiber.StatusOK, "Successful get", response)
}

func (r *serviceRoutes) servicesByUserID(ctx *fiber.Ctx) error {
	jwtData := ctx.Locals("jwt").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	userID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "http - v1 - servicesByUserID")
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	services, err := r.s.GetByUserID(userID)
	if err != nil {
		r.l.Error(err, "http - v1 - servicesByUserID")
		if err == usecase.ErrUserNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "User with provided id not found", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	respServices := make([]serviceResponse, len(services), cap(services))

	for i, v := range services {
		temp := serviceResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Price:       v.Price,
			IsClosed:    v.IsClosed,
		}
		respServices[i] = temp
	}

	return successResponse(ctx, fiber.StatusOK, "Successful get services", respServices)
}

func (r *serviceRoutes) deleteService(ctx *fiber.Ctx) error {
	request := doDeleteServiceRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.l.Error(err, "http - v1 - service - delete")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(&request)
	if !ok {
		r.l.Error(errs, "http - v1 - service - delete")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	serviceID, err := uuid.Parse(request.ID)
	if err != nil {
		r.l.Error(err, "http - v1 - service - delete")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid uuid", err)
	}

	delService, err := r.s.Delete(serviceID)
	if err != nil {
		r.l.Error(err, "http - v1 - service - delete")
		if err == usecase.ErrServiceNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "Service with provided id not exists", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := serviceResponse{
		ID:          delService.ID,
		Title:       delService.Title,
		Description: delService.Description,
		Price:       delService.Price,
		IsClosed:    delService.IsClosed,
	}

	return successResponse(ctx, fiber.StatusOK, "Successful delete", response)
}
