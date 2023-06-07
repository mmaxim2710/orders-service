package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/controller/http/v1/middleware"
	"github.com/mmaxim2710/orders-service/internal/usecase"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"github.com/mmaxim2710/orders-service/pkg/validations"
)

type orderRoutes struct {
	order  usecase.Order
	logger logger.Interface
}

func newOrderRoutes(handler fiber.Router, order usecase.Order, logger logger.Interface) {
	r := &orderRoutes{
		order:  order,
		logger: logger,
	}

	h := handler.Group("/orders")
	{
		h.Post("", middleware.Protected(), r.create)
		h.Get("/:orderID", middleware.Protected(), r.getByID)
	}
}

type (
	doCreateOrderRequest struct {
		ServiceIDs []string `json:"service_ids" validate:"required"`
	}

	serviceEntity struct {
		ServiceID uuid.UUID `json:"service_id"`
		Status    string    `json:"status"`
	}

	orderResponse struct {
		OrderID  uuid.UUID       `json:"order_id"`
		Services []serviceEntity `json:"services"`
	}
)

// @Summary     Create Order
// @Description Create new order
// @ID          create
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Param       request body doCreateOrderRequest true "List of services for order"
// @Success     200 {object} Response{data=orderResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /orders [post]
func (r *orderRoutes) create(ctx *fiber.Ctx) error {
	request := doCreateOrderRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.logger.Error(err, "http - v1 - order - create")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(&request)
	if !ok {
		r.logger.Error(errs, "http - v1 - order - create")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	serviceUUIDs := make([]uuid.UUID, 0, len(request.ServiceIDs))
	for _, v := range request.ServiceIDs {
		id, err := uuid.Parse(v)
		if err != nil {
			r.logger.Error(err, "http - v1 - order - create")
			return errorResponse(ctx, fiber.StatusBadRequest, "Error while parsing service IDs", err)
		}
		serviceUUIDs = append(serviceUUIDs, id)
	}

	newOrders, err := r.order.Create(serviceUUIDs)
	if err != nil {
		r.logger.Error(err, "http - v1 - order - create")
		if err == usecase.ErrServiceNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "Service not found", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	if len(newOrders) == 0 {
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", ErrInternalServer)
	}

	response := orderResponse{}
	response.Services = make([]serviceEntity, 0, len(newOrders))
	response.OrderID = newOrders[0].OrderID
	for _, v := range newOrders {
		response.Services = append(response.Services, serviceEntity{
			ServiceID: v.ServiceID,
			Status:    v.Status,
		})
	}

	return successResponse(ctx, fiber.StatusOK, "Success create order", response)
}

// @Summary     Get Order by id
// @Description Get existing order by UUID
// @ID          getByID
// @Tags  	    order
// @Accept      */*
// @Produce     json
// @Param       orderID path string true "Order data"
// @Success     200 {object} Response{data=orderResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /orders/{orderID} [get]
func (r *orderRoutes) getByID(ctx *fiber.Ctx) error {
	strID := ctx.Params("orderID")

	orderID, err := uuid.Parse(strID)
	if err != nil {
		r.logger.Error(err, "http - v1 - order - getByID")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid order uuid", err)
	}

	order, err := r.order.GetByOrderID(orderID)
	if err != nil {
		r.logger.Error(err, "http - v1 - order - getByID")
		if err == usecase.ErrOrderNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "Orders with provided id not found", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := orderResponse{}
	response.Services = make([]serviceEntity, 0, len(order))
	response.OrderID = order[0].OrderID
	for _, v := range order {
		response.Services = append(response.Services, serviceEntity{
			ServiceID: v.ServiceID,
			Status:    v.Status,
		})
	}

	return successResponse(ctx, fiber.StatusOK, "Successful get", response)
}
