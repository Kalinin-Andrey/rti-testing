package product

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
)

// Product entity
type Product struct {
	ID				uint					`gorm:"PRIMARY_KEY" json:"id"`
	Name			string					`gorm:"type:varchar(100)" json:"name"`
	Components		[]component.Component	`gorm:"FOREIGNKEY:ProductID;association_autoupdate:false" json:"components"`

	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time				`sql:"index"`
}

const (
	// TableName const
	TableName	= "product"
)

// Validate func
func (e Product) Validate() error {

	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(2, 100), is.PrintableASCII),
	)
}

// TableName func
func (e Product) TableName() string {
	return TableName
}

// New func is a constructor
func New() *Product {
	return &Product{}
}

func (e Product) IsValid() bool {
	var isThereAMainComponent bool
	for _, c := range e.Components {
		if c.IsMain {
			isThereAMainComponent = true
		}
	}
	return isThereAMainComponent
}

func (e Product) TotalCost() *price.Price {
	pr := &price.Price{
		Type:	price.TypeCost,
	}
	for _, c := range e.Components {
		p := c.TotalCost()
		pr.Cost += p.Cost
	}
	return pr
}