package mock

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
)

// UserRepository is a mock for ProductRepository
type ProductRepository struct {
	Response struct {
		Get		struct {
			Entity	*product.Product
			Err		error
		}
		First	struct {
			Entity	*product.Product
			Err		error
		}
		Query	struct {
			List	[]product.Product
			Err		error
		}
		Create	struct {
			Entity	*product.Product
			Err		error
		}
		Update	struct {
			Entity	*product.Product
			Err		error
		}
		Delete	struct {
			Err		error
		}
	}
}

var _ product.IRepository = (*ProductRepository)(nil)

func (r ProductRepository) SetDefaultConditions(conditions map[string]interface{}) {}

func (r ProductRepository) Get(ctx context.Context, id uint) (*product.Product, error) {
	return r.Response.Get.Entity, r.Response.Get.Err
}

func (r ProductRepository) First(ctx context.Context, user *product.Product) (*product.Product, error) {
	return r.Response.First.Entity, r.Response.First.Err
}

func (r ProductRepository) Query(ctx context.Context, offset, limit uint) ([]product.Product, error) {
	return r.Response.Query.List, r.Response.Query.Err
}

func (r ProductRepository) Create(ctx context.Context, entity *product.Product) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Create.Err
}

func (r ProductRepository) Update(ctx context.Context, entity *product.Product) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Update.Err
}

func (r ProductRepository) Delete(ctx context.Context, id uint) error {
	return r.Response.Delete.Err
}
