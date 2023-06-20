package providers

import (
	"goravel/app/policies"

	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type AuthServiceProvider struct {
}

func (receiver *AuthServiceProvider) Register(app foundation.Application) {}

func (receiver *AuthServiceProvider) Boot(app foundation.Application) {

	// Polcies definitions
	facades.Gate().Define("update-todo", policies.NewTodoPolicy().Update)
	facades.Gate().Define("destroy-todo", policies.NewTodoPolicy().Destroy)

}
