package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdatedTodoRequest struct {
	Title string `form:"title" json:"title"`
	// CompletedAt *time.Time `form:"completed_at" json:"completed_at,omitempty"`
}

func (r *UpdatedTodoRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdatedTodoRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"title": "required|max_len:255",
		// "completed_at": "date",
	}
}

func (r *UpdatedTodoRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatedTodoRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdatedTodoRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
