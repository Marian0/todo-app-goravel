package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Welcome to": facades.Config.Env("APP_NAME"),
		})
	})

	authController := controllers.NewAuthController()
	facades.Route.Post("/auth/register", authController.Register)
	facades.Route.Post("/auth/login", authController.Login)
}
