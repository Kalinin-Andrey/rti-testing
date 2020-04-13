package offer

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
)

// Offer entity
type Offer struct {
	product.Product
	TotalCost	price.Price	`json:"totalCost"`
}

// New func is a constructor
func New() *Offer {
	return &Offer{}
}

