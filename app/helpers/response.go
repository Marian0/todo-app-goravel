package helpers

import (
	"github.com/goravel/framework/contracts/http"
)

func RespondSuccess(ctx http.Context, statusCode int, data any) {
	ctx.Response().Json(statusCode, http.Json{
		"data": data,
	})
}

func RespondError(ctx http.Context, statusCode int, errs any) {
	ctx.Response().Json(statusCode, http.Json{
		"errors": errs,
	})
}
