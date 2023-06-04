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

	// hash users password
	hashedPassword, err := facades.Hash.Make(user.Password)
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
		})
		return
	}
	user.Password = hashedPassword

	if err := facades.Orm.Query().Create(&user); err != nil {
		//@todo: implement proper error handler to hide db errors
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": err.Error(),
		})
		return
	}

	token, err := facades.Auth.LoginUsingID(ctx, user.ID)
	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
		})
		return
	}

	//@todo: implement proper DTO approach to transform models into JSON reponse
	ctx.Response().Success().Json(http.Json{
		"ID":    user.ID,
		"name":  user.Name,
		"token": token,
	})
}
