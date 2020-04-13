package mock

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
)

// UserRepository is a mock for ComponentRepository
type ComponentRepository struct {
	Response struct {
		Get		struct {
			Entity	*component.Component
			Err		error
		}
		First	struct {
			Entity	*component.Component
			Err		error
		}
		Query	struct {
			List	[]component.Component
			Err		error
		}
		Create	struct {
			Entity	*component.Component
			Err		error
		}
		Update	struct {
			Entity	*component.Component
			Err		error
		}
		Delete	struct {
			Err		error
		}
	}
}

var _ component.IRepository = (*ComponentRepository)(nil)

func (r ComponentRepository) SetDefaultConditions(conditions map[string]interface{}) {}

func (r ComponentRepository) Get(ctx context.Context, id uint) (*component.Component, error) {
	return r.Response.Get.Entity, r.Response.Get.Err
}

func (r ComponentRepository) First(ctx context.Context, user *component.Component) (*component.Component, error) {
	return r.Response.First.Entity, r.Response.First.Err
}

func (r ComponentRepository) Query(ctx context.Context, offset, limit uint) ([]component.Component, error) {
	return r.Response.Query.List, r.Response.Query.Err
}

func (r ComponentRepository) Create(ctx context.Context, entity *component.Component) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Create.Err
}

func (r ComponentRepository) Update(ctx context.Context, entity *component.Component) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Update.Err
}

func (r ComponentRepository) Delete(ctx context.Context, id uint) error {
	return r.Response.Delete.Err
}
