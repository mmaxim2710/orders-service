package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmaxim2710/orders-service/internal/usecase"
)

type userRoutes struct {
	u usecase.User
}

func newUserRoutes(handler fiber.Router, u usecase.User) {
	r := &userRoutes{
		u: u,
	}

	h := handler.Group("/user")
	{
		h.Post("/register", r.registerUser)
	}
}

type doRegisterUserRequest struct {
	Login     string `json:"login,omitempty" example:"testUser"`
	Email     string `json:"email,omitempty" example:"user@example.com"`
	FirstName string `json:"first_name,omitempty" example:"Ivan"`
	LastName  string `json:"last_name,omitempty" example:"Pupkin"`
	Password  string `json:"password,omitempty" example:"supersecretpassword"`
}

type registerUserResponse struct {
	Login     string `json:"login,omitempty" example:"testUser"`
	Email     string `json:"email,omitempty" example:"user@example.com"`
	FirstName string `json:"first_name,omitempty" example:"Ivan"`
	LastName  string `json:"last_name,omitempty" example:"Pupkin"`
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
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	user, err := r.u.RegisterUser(request.Login, request.Email, request.FirstName, request.LastName, request.Password)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError, "Error while register user", err)
	}

	response := &registerUserResponse{
		Login:     user.Login,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return successResponse(ctx, fiber.StatusOK, "Successfully registered user", response)
}
