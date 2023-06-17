package policies

import (
	"context"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/auth/access"
)

type TodoPolicy struct {
}

func NewTodoPolicy() *TodoPolicy {
	return &TodoPolicy{}
}

func (r *TodoPolicy) Update(ctx context.Context, arguments map[string]any) access.Response {
	user := ctx.Value("user").(models.User)
	todo := arguments["todo"].(models.Todo)

	//@todo: check panics while object casting

	if user.ID == todo.UserID {
		return access.NewAllowResponse()
	} else {
		return access.NewDenyResponse("You do not own this todo")
	}
}
