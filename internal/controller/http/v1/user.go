package v1

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mmaxim2710/orders-service/internal/usecase"
	"github.com/mmaxim2710/orders-service/pkg/logger"
	"github.com/mmaxim2710/orders-service/pkg/validations"
)

type userRoutes struct {
	u usecase.User
	l logger.Interface
}

func newUserRoutes(handler fiber.Router, u usecase.User, l logger.Interface) {
	r := &userRoutes{
		u: u,
		l: l,
	}

	h := handler.Group("/user")
	{
		h.Post("/register", r.registerUser)
	}
}

type doRegisterUserRequest struct {
	Login     string `json:"login" example:"testUser" validate:"required,min=3,max=255"`
	Email     string `json:"email" example:"user@example.com" validate:"required,email"`
	FirstName string `json:"first_name" example:"Ivan" validate:"required,min=2,max=128"`
	LastName  string `json:"last_name" example:"Pupkin" validate:"required,min=2,max=128"`
	Password  string `json:"password" example:"supersecretpassword" validate:"required,min=8,max=64"`
}

type registerUserResponse struct {
	ID        uuid.UUID `json:"id,omitempty" example:"89db2ce2-f2c6-4d59-a014-8b68d19b783c"`
	Login     string    `json:"login,omitempty" example:"testUser"`
	Email     string    `json:"email,omitempty" example:"user@example.com"`
	FirstName string    `json:"first_name,omitempty" example:"Ivan"`
	LastName  string    `json:"last_name,omitempty" example:"Pupkin"`
}

// @Summary     Register user
// @Description Register a new user with passed params
// @ID          register-user
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body doRegisterUserRequest true "Set up user"
// @Success     200 {object} registerUserResponse
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/register [post]
func (r *userRoutes) registerUser(ctx *fiber.Ctx) error {
	request := &doRegisterUserRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.l.Error(err, "http - v1 - registerUser")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(request)
	if !ok {
		r.l.Error(errors.New("validation failed"), "http - v1 - registerUser")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	user, err := r.u.RegisterUser(request.Login, request.Email, request.FirstName, request.LastName, request.Password)
	if err != nil {
		r.l.Error(err, "")
		return errorResponse(ctx, fiber.StatusInternalServerError, "Error while register user", err)
	}

	response := &registerUserResponse{
		ID:        user.ID,
		Login:     user.Login,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return successResponse(ctx, fiber.StatusOK, "Successfully registered user", response)
}
