package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Welcome to": facades.Config.Env("APP_NAME"),
		})
	})

	authController := controllers.NewAuthController()

	// Public endpoints
	facades.Route.Post("/auth/register", authController.Register)
	facades.Route.Post("/auth/login", authController.Login)

	// Secuged endpoints
	facades.Route.Middleware(middleware.Jwt()).Group(func(route route.Route) {
		route.Get("/auth/me", authController.Me)
	})
}
