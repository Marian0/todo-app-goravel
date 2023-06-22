package controllers

import (
	"goravel/app/helpers"
	"goravel/app/http/dtos"
	"goravel/app/http/requests/auth_requests"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	//Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		//Inject services
	}
}

// POST /auth/register
// Registers a new user
func (r *AuthController) Register(ctx http.Context) {
	var registerRequest auth_requests.RegisterRequest
	errors, err := ctx.Request().ValidateRequest(&registerRequest)
	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if errors != nil {
		helpers.RespondError(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}

	// hash users password
	hashedPassword, err := facades.Hash().Make(registerRequest.Password)
	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var user models.User
	user.Name = registerRequest.Name
	user.Email = registerRequest.Email
	user.Password = hashedPassword

	if err := facades.Orm().Query().Create(&user); err != nil {
		//@todo: implement proper error handler to hide db errors
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := facades.Auth().LoginUsingID(ctx, user.ID)
	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondSuccess(ctx, http.StatusOK, http.Json{
		"user":  dtos.UserToDTO(user),
		"token": token,
	})
}

// POST /auth/login
// Check user creds and return jwt
func (r *AuthController) Login(ctx http.Context) {
	var loginRequest auth_requests.LoginRequest
	errors, err := ctx.Request().ValidateRequest(&loginRequest)
	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if errors != nil {
		helpers.RespondError(ctx, http.StatusUnprocessableEntity, errors.All())
		return
	}

	// Look for user
	var user models.User
	err = facades.Orm().Query().Where("email = ?", loginRequest.Email).FindOrFail(&user)
	if err != nil {
		helpers.RespondError(ctx, http.StatusBadRequest, "Bad username or password")
		return
	}

	//password check
	if !facades.Hash().Check(loginRequest.Password, user.Password) {
		helpers.RespondError(ctx, http.StatusBadRequest, "Bad username or password")
		return
	}

	token, err := facades.Auth().LoginUsingID(ctx, user.ID)
	if err != nil {
		helpers.RespondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondSuccess(ctx, http.StatusOK, http.Json{
		"user":  dtos.UserToDTO(user),
		"token": token,
	})
}

// GET auth/me
// returns current user by processing the Bearer token with Jwt middleware
func (r *AuthController) Me(ctx http.Context) {
	user := ctx.Value("user").(models.User)
	helpers.RespondSuccess(ctx, http.StatusOK, dtos.UserToDTO(user))
}
