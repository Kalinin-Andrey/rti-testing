package price

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"
)

// Price entity
type Price struct {
	ID                  uint                					`gorm:"PRIMARY_KEY" json:"id"`
	ComponentID			uint									`sql:"type:int REFERENCES component(id)" json:"componentId"`
	Cost                float64             					`json:"cost"`
	Type                string              					`json:"priceType,omitempty"`
	RuleApplicabilities []ruleapplicability.RuleApplicability	`gorm:"FOREIGNKEY:PriceID;association_autoupdate:false" json:"ruleApplicabilities,omitempty"`

	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			*time.Time								`sql:"index"`
}

const (
	// TableName const
	TableName		= "price"
	// TypeCost const
	TypeCost		= "COST"
	// TypeDiscount const
	TypeDiscount	= "Discount"
)

var Types []interface{} = []interface{}{
	TypeCost,
	TypeDiscount,
}

// Validate func
func (e Price) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Type, validation.Required, validation.Length(2, 100), is.Alpha, validation.In(Types...)),
	)
}

// TableName func
func (e Price) TableName() string {
	return TableName
}

// New func is a constructor
func New() *Price {
	return &Price{}
}


func (e Price) IsValid() bool {
	return len(e.RuleApplicabilities) > 0
}
