package middleware

import (
	"errors"
	"goravel/app/helpers"
	"goravel/app/models"
	"net/http"

	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Jwt() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, helpers.GetJSONErrorBody("unauthorized"))
			return
		}

		if _, err := facades.Auth().Parse(ctx, token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				token, err = facades.Auth().Refresh(ctx)
				if err != nil {
					// Refresh time exceeded
					ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, helpers.GetJSONErrorBody("unauthorized"))
					return
				}

				token = "Bearer " + token
			} else {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, helpers.GetJSONErrorBody("unauthorized"))
				return
			}
		}

		// Query DB and validates USER to be injected into the context
		var user models.User
		if err := facades.Auth().User(ctx, &user); err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, helpers.GetJSONErrorBody("unauthorized"))
			return
		}
		ctx.WithValue("user", user)

		//
		ctx.Response().Header("Authorization", token)
		ctx.Request().Next()
	}
}
