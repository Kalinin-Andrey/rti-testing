package ruleapplicability

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"strconv"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

)

// RuleApplicability entity
type RuleApplicability struct {
	ID				uint					`gorm:"PRIMARY_KEY" json:"id"`
	PriceID			uint					`sql:"type:int REFERENCES price(id)" json:"priceId"`
	CodeName		string					`gorm:"type:varchar(100)" json:"codeName"`
	Operator		string					`gorm:"type:varchar(100)" json:"operator"`
	Value			string					`gorm:"type:varchar(100)" json:"value"`

	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time				`sql:"index"`
}

const (
	// TableName const
	TableName					= "rule_applicability"
	OperatorEqual              = "EQ"
	OperatorGreaterThanOrEqual = "GTE"
	OperatorLessThanOrEqual    = "LTE"
)

var Operators []interface{} = []interface{}{
	OperatorEqual,
	OperatorGreaterThanOrEqual,
	OperatorLessThanOrEqual,
}

// Validate func
func (e RuleApplicability) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.CodeName, validation.Required, validation.Length(2, 100), is.Alpha),
		validation.Field(&e.Operator, validation.Required, validation.Length(2, 100), is.Alpha, validation.In(Operators...)),
	)
}

// TableName func
func (e RuleApplicability) TableName() string {
	return TableName
}

// New func is a constructor
func New() *RuleApplicability {
	return &RuleApplicability{}
}

func (e RuleApplicability) IsSatisfy(c condition.Condition) (isSatisfy bool, err error) {

	if c.RuleName == e.CodeName {
		switch e.Operator {
		case OperatorEqual:
			isSatisfy = c.Value == e.Value
		case OperatorGreaterThanOrEqual:
			cVal, err := strconv.ParseFloat(c.Value, 64)
			if err != nil {
				return isSatisfy, err
			}
			rVal, err := strconv.ParseFloat(e.Value, 64)
			if err != nil {
				return isSatisfy, err
			}
			isSatisfy = cVal >= rVal
		case OperatorLessThanOrEqual:
			cVal, err := strconv.ParseFloat(c.Value, 64)
			if err != nil {
				return isSatisfy, err
			}
			rVal, err := strconv.ParseFloat(e.Value, 64)
			if err != nil {
				return isSatisfy, err
			}
			isSatisfy = cVal <= rVal
		}
	}

	return isSatisfy, nil
}