package component

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
)

// Component entity
type Component struct {
	ID				uint					`gorm:"PRIMARY_KEY" json:"id"`
	ProductID		uint					`sql:"type:int REFERENCES product(id)" json:"productId"`
	IsMain			bool					`json:"isMain,omitempty"`
	Name			string					`gorm:"type:varchar(100)" json:"name"`
	Prices			[]price.Price			`gorm:"FOREIGNKEY:ComponentID;association_autoupdate:false" json:"prices"`

	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time				`sql:"index"`
}

const (
	// TableName const
	TableName	= "component"
)

// Validate func
func (e Component) Validate() error {

	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(2, 100), is.PrintableASCII),
	)
}

// TableName func
func (e Component) TableName() string {
	return TableName
}

// New func is a constructor
func New() *Component {
	return &Component{}
}


func (e Component) IsValid() bool {
	var typeCosts uint
	for _, p := range e.Prices {
		if p.Type == price.TypeCost {
			typeCosts++
		}
	}
	return typeCosts == 1
}

func (e Component) TotalCost() *price.Price {
	var cost, discount float64
	pr := &price.Price{
		Type:	price.TypeCost,
	}
	for _, p := range e.Prices {
		switch p.Type {
		case price.TypeCost:
			cost = p.Cost
		case price.TypeDiscount:
			if p.Cost > discount {
				discount = p.Cost
			}
		}
	}
	if discount > 0 {
		cost -= cost * discount / 100
	}
	pr.Cost = cost
	return pr
}

