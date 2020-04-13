package condition

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

)

// Condition value object
type Condition struct {
	RuleName	string	`json:"ruleName"`
	Value		string	`json:"value"`
}

// Validate func
func (e Condition) Validate() error {

	return validation.ValidateStruct(&e,
		validation.Field(&e.RuleName, validation.Required, validation.Length(2, 100), is.Alpha),
		validation.Field(&e.Value, validation.Required, validation.Length(2, 100), is.PrintableASCII),
	)
}

// New func is a constructor
func New() *Condition {
	return &Condition{}
}

