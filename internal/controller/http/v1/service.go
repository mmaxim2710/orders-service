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
	}
}

type doCreateServiceRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type createServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

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
