package offer

import (
	"context"
	"reflect"

	"github.com/Kalinin-Andrey/rti-testing/pkg/log"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
)

// IService encapsulates usecase logic for offer.
type IService interface {
	NewEntity() *Offer
	Calculate(product *product.Product, conditions []condition.Condition) (offer *Offer, err error)
	CalculateByProductID(ctx context.Context, productId uint, conditions []condition.Condition) (offer *Offer, err error)
}

type service struct {
	//Domain     Domain
	productService	product.IService
	logger			log.ILogger
}

// NewService creates a new service.
func NewService(productService product.IService, logger log.ILogger) IService {
	return &service{productService, logger}
}

// Defaults returns defaults params
func (s service) defaultConditions() map[string]interface{} {
	return map[string]interface{}{
	}
}

func (s service) NewEntity() *Offer {
	return &Offer{}
}

func (s service) Calculate(prod *product.Product, conditions []condition.Condition) (offer *Offer, err error) {

	if prod == nil || conditions == nil {
		return nil, nil
	}
	emptyProduct := &product.Product{}

	if reflect.DeepEqual(*prod, *emptyProduct) || len(conditions) == 0 {
		return &Offer{}, nil
	}

	p, err := s.productService.BuildByConditons(prod, conditions)
	if err != nil {
		if err == apperror.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &Offer{
		Product:	*p,
		TotalCost:	*p.TotalCost(),
	}, nil
}

func (s service) CalculateByProductID(ctx context.Context, productId uint, conditions []condition.Condition) (offer *Offer, err error) {
	prod, err := s.productService.Get(ctx, productId)
	if err != nil {
		return nil, err
	}
	return s.Calculate(prod, conditions)
}
