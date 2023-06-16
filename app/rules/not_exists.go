package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

/**
 * not_exists validator rule
 * not_exists verify a value does not exist in a table field
 * Usage: not_exists:table_name,field_name,field_name,field_name
 * Example: not_exists:users,phone,email
 */

type NotExists struct {
}

// Signature The name of the rule.
func (receiver *NotExists) Signature() string {
	return "not_exists"
}

// Passes Determine if the validation rule passes.
func (receiver *NotExists) Passes(_ validation.Data, val any, options ...any) bool {

	tableName := options[0].(string)
	fieldName := options[1].(string)
	requestValue, ok := val.(string)
	// check if requestValue could be transform into a string
	if !ok {
		return true
	}

	if len(requestValue) == 0 {
		return false
	}

	var count int64
	query := facades.Orm.Query().Table(tableName).Where(fieldName, requestValue)
	if len(options) > 2 {
		for i := 2; i < len(options); i++ {
			query = query.OrWhere(options[i].(string), requestValue)
		}
	}
	err := query.Count(&count)
	if err != nil {
		return false
	}

	return count == 0
}

// Message Get the validation error message.
func (receiver *NotExists) Message() string {
	return "record already exists"
}
