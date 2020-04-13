package mock

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
)

// UserRepository is a mock for UserRepository
type PriceRepository struct {
	Response struct {
		Get		struct {
			Entity	*price.Price
			Err		error
		}
		First	struct {
			Entity	*price.Price
			Err		error
		}
		Query	struct {
			List	[]price.Price
			Err		error
		}
		Create	struct {
			Entity	*price.Price
			Err		error
		}
		Update	struct {
			Entity	*price.Price
			Err		error
		}
		Delete	struct {
			Err		error
		}
	}
}

var _ price.IRepository = (*PriceRepository)(nil)

func (r PriceRepository) SetDefaultConditions(conditions map[string]interface{}) {}

func (r PriceRepository) Get(ctx context.Context, id uint) (*price.Price, error) {
	return r.Response.Get.Entity, r.Response.Get.Err
}

func (r PriceRepository) First(ctx context.Context, user *price.Price) (*price.Price, error) {
	return r.Response.First.Entity, r.Response.First.Err
}

func (r PriceRepository) Query(ctx context.Context, offset, limit uint) ([]price.Price, error) {
	return r.Response.Query.List, r.Response.Query.Err
}

func (r PriceRepository) Create(ctx context.Context, entity *price.Price) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Create.Err
}

func (r PriceRepository) Update(ctx context.Context, entity *price.Price) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Update.Err
}

func (r PriceRepository) Delete(ctx context.Context, id uint) error {
	return r.Response.Delete.Err
}
