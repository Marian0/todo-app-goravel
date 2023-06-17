package providers

import (
	"goravel/app/policies"

	"github.com/goravel/framework/facades"
)

type AuthServiceProvider struct {
}

func (receiver *AuthServiceProvider) Register() {

}

func (receiver *AuthServiceProvider) Boot() {

	// Polcies definitions
	facades.Gate.Define("update-todo", policies.NewTodoPolicy().Update)

}
