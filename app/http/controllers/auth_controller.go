package controllers

import (
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

func (r *AuthController) Register(ctx http.Context) {
	validator, err := ctx.Request().Validate(map[string]string{
		"name":     "required",
		"email":    "required|email|not_exists:users,email",
		"password": "required|min:5",
	})
	if err != nil {
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
		return
	}
	if validator.Fails() {
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": validator.Errors().All(),
		})
		return
	}

	var user models.User
	if err := validator.Bind(&user); err != nil {
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
		return
	}

	if err := facades.Orm.Query().Create(&models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		//@todo: implement proper error handler to hide db errors
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

	//@todo: implement proper DTO approach to transform models into JSON reponse
	ctx.Response().Success().Json(http.Json{
		"ID":   user.ID,
		"name": user.Name,
	})
}
