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

type userRoutes struct {
	user   usecase.User
	logger logger.Interface
}

func newUserRoutes(handler fiber.Router, u usecase.User, l logger.Interface) {
	r := &userRoutes{
		user:   u,
		logger: l,
	}

	h := handler.Group("/user")
	{
		h.Post("/register", r.registerUser)
		h.Post("/login", r.login)
		h.Patch("/refresh", r.refresh)
		h.Patch("/update", middleware.Protected(), r.update)
		h.Delete("/:userID", middleware.Protected(), r.delete)
	}
}

type (
	doRegisterUserRequest struct {
		Login     string `json:"login" example:"testUser" validate:"required,min=3,max=255"`
		Email     string `json:"email" example:"user@example.com" validate:"required,email"`
		FirstName string `json:"first_name" example:"Ivan" validate:"required,min=2,max=128"`
		LastName  string `json:"last_name" example:"Pupkin" validate:"required,min=2,max=128"`
		Password  string `json:"password" example:"supersecretpassword" validate:"required,min=8,max=64"`
	}
	doLoginRequest struct {
		Email    string `json:"email" example:"user@example.com" validate:"required,email"`
		Password string `json:"password" example:"supersecretpassword" validate:"required,min=8,max=64"`
	}
	doRefreshRequest struct {
		UserID       string `json:"user_id" example:"89db2ce2-f2c6-4d59-a014-8b68d19b783c" validate:"required,uuid4"`
		RefreshToken string `json:"refresh_token" example:"fsd7-1fsdfs23SDFfsdf3" validate:"required"`
	}
	doUpdateRequest struct {
		Email     string `json:"email" example:"user@example.com" validation:"required,email"`
		FirstName string `json:"first_name" example:"Vasya" validation:"required,min=2,max=128"`
		LastName  string `json:"last_name" example:"Pupkin" validation:"required,min=2,max=128"`
	}

	userResponse struct {
		ID        uuid.UUID `json:"id" example:"89db2ce2-f2c6-4d59-a014-8b68d19b783c"`
		Login     string    `json:"login" example:"testUser"`
		Email     string    `json:"email" example:"user@example.com"`
		FirstName string    `json:"first_name" example:"Ivan"`
		LastName  string    `json:"last_name" example:"Pupkin"`
	}
)

// @Summary     Register user
// @Description Register a new user with passed params
// @ID          register-user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body doRegisterUserRequest true "Set up user"
// @Success     200 {object} Response{data=userResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/register [post]
func (r *userRoutes) registerUser(ctx *fiber.Ctx) error {
	request := doRegisterUserRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.logger.Error(err, "http - v1 - registerUser")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(request)
	if !ok {
		r.logger.Error(ErrValidationFailed, "http - v1 - registerUser")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	user, err := r.user.RegisterUser(request.Login, request.Email, request.FirstName, request.LastName, request.Password)
	if err != nil {
		r.logger.Error(err)
		if err == usecase.ErrUserExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "User with provided email is exists", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Error while register user", err)
	}

	response := &userResponse{
		ID:        user.ID,
		Login:     user.Login,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return successResponse(ctx, fiber.StatusOK, "Successfully registered user", response)
}

// @Summary     Login
// @Description Login to user account
// @ID          login
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body doLoginRequest true "JWT token, refresh token"
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/login [post]
func (r *userRoutes) login(ctx *fiber.Ctx) error {
	request := doLoginRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(request)
	if !ok {
		r.logger.Error(ErrValidationFailed, "http - v1 - login")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	result, err := r.user.Login(request.Email, request.Password)
	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		return errorResponse(ctx, fiber.StatusBadRequest, "Login failed", err)
	}

	return successResponse(ctx, fiber.StatusOK, "Login success", result)
}

// @Summary     Refresh
// @Description Refresh JWT token by refresh token
// @ID          refresh
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body doRefreshRequest true "New JWT token"
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/refresh [patch]
func (r *userRoutes) refresh(ctx *fiber.Ctx) error {
	request := doRefreshRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.logger.Error(err, "http - v1 - refresh")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(request)
	if !ok {
		r.logger.Error(ErrValidationFailed, "http - v1 - refresh")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	result, err := r.user.Refresh(request.RefreshToken, request.UserID)
	if err != nil {
		r.logger.Error(err, "http - v1 - refresh")
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}
	return successResponse(ctx, fiber.StatusOK, "Successful refresh", result)
}

// @Summary     Update
// @Description Update users info
// @ID          update
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body doUpdateRequest true "User data"
// @Success     200 {object} Response{data=userResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/update [patch]
func (r *userRoutes) update(ctx *fiber.Ctx) error {
	request := doUpdateRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		r.logger.Error(err, "http - v1 - update")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}

	ok, errs := validations.UniversalValidation(request)
	if !ok {
		r.logger.Error(ErrValidationFailed, "http - v1 - update")
		return errorResponse(ctx, fiber.StatusBadRequest, "Validation failed", errs)
	}

	jwtData := ctx.Locals("jwt").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	userID, err := uuid.Parse(id)
	if err != nil {
		r.logger.Error(err, "http - v1 - update")
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	newUser, err := r.user.Update(userID, request.Email, request.FirstName, request.LastName)
	if err != nil {
		r.logger.Error(err, "http - v1 - refresh")
		if err == usecase.ErrUserExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "User with provided email is exists", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := &userResponse{
		ID:        newUser.ID,
		Login:     newUser.Login,
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
	}
	return successResponse(ctx, fiber.StatusOK, "Successful update", response)
}

// @Summary     Delete
// @Description Delete user
// @ID          delete
// @Tags  	    user
// @Accept      */*
// @Produce     json
// @Param       userID path string true "User data"
// @Success     200 {object} Response{data=userResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /user/{userID} [delete]
func (r *userRoutes) delete(ctx *fiber.Ctx) error {
	userID := ctx.Params("userID")
	jwtData := ctx.Locals("jwt").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	if userID != id {
		return errorResponse(ctx, fiber.StatusBadRequest, "Trying to delete user that is not associated wit token", ErrEntitiesMismatch)
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		r.logger.Error(err, "http - v1 - user - delete")
		return errorResponse(ctx, fiber.StatusBadRequest, "Invalid uuid in request params", err)
	}

	delUser, err := r.user.Delete(userUUID)
	if err != nil {
		r.logger.Error(err, "http - v1 - user - delete")
		if err == usecase.ErrUserNotExists {
			return errorResponse(ctx, fiber.StatusBadRequest, "user associated with provided token not exists", err)
		}
		if err == usecase.ErrUserHasNonClosedServices {
			return errorResponse(ctx, fiber.StatusBadRequest, "user has non closed services", err)
		}
		return errorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err)
	}

	response := userResponse{
		ID:        delUser.ID,
		Login:     delUser.Login,
		Email:     delUser.Email,
		FirstName: delUser.FirstName,
		LastName:  delUser.LastName,
	}

	return successResponse(ctx, fiber.StatusOK, "Successful delete", response)
}
