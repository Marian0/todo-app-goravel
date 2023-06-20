package helpers

import (
	"errors"

	"github.com/goravel/framework/facades"
)

func ValidateUUID(value any) (*string, error) {
	uuid, ok := value.(string)
	// check if requestValue could be transform into a string
	if !ok {
		return nil, errors.New("uuid not valid type")
	}

	// use validator to check uuid
	validator, err := facades.Validation().Make(map[string]any{
		"uuid": uuid,
	}, map[string]string{
		"uuid": "required|uuid",
	})

	if err != nil {
		return nil, errors.New("uuid not valid type")
	}

	if validator.Fails() {
		return nil, errors.New("invalid uuid type")
	}

	return &uuid, nil

}
