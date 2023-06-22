package helpers

import (
	"github.com/goravel/framework/contracts/http"
)

// Body formatters
func GetJSONSuccessBody(errs any) http.Json {
	return http.Json{
		"errors": errs,
	}
}

func GetJSONErrorBody(errs any) http.Json {
	return http.Json{
		"errors": errs,
	}
}

// Response helpers
func RespondSuccess(ctx http.Context, statusCode int, data any) {
	ctx.Response().Json(statusCode, GetJSONSuccessBody(data))
}

func RespondError(ctx http.Context, statusCode int, errors any) {
	ctx.Response().Json(statusCode, GetJSONErrorBody(errors))
}
