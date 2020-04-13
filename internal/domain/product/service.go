package product

import (
	"context"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"

	"github.com/pkg/errors"

	"github.com/Kalinin-Andrey/rti-testing/pkg/log"
)

const MaxLIstLimit = 1000

// IService encapsulates usecase logic for user.
type IService interface {
	NewEntity() *Product
	Get(ctx context.Context, id uint) (*Product, error)
	First(ctx context.Context, entity *Product) (*Product, error)
	Query(ctx context.Context, offset, limit uint) ([]Product, error)
	List(ctx context.Context) ([]Product, error)
	//Count(ctx context.Context) (uint, error)
	Create(ctx context.Context, entity *Product) error
	Update(ctx context.Context, entity *Product) error
	Delete(ctx context.Context, id uint) (error)
	BuildByConditons(product *Product, conditions []condition.Condition) (*Product, error)
}

type service struct {
	//Domain     Domain
	repo       			IRepository
	componentService	component.IService
	logger     			log.ILogger
}

// NewService creates a new service.
func NewService(repo IRepository, componentService component.IService, logger log.ILogger) IService {
	s := &service{repo, componentService, logger}
	repo.SetDefaultConditions(s.defaultConditions())
	return s
}

// Defaults returns defaults params
func (s service) defaultConditions() map[string]interface{} {
	return map[string]interface{}{
	}
}

func (s service) NewEntity() *Product {
	return &Product{}
}

// Get returns the entity with the specified ID.
func (s service) Get(ctx context.Context, id uint) (*Product, error) {
	entity, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not get a product by id: %v", id)
	}
	return entity, nil
}

func (s service) First(ctx context.Context, entity *Product) (*Product, error) {
	return s.repo.First(ctx, entity)
}

/*
// Count returns the number of items.
func (s service) Count(ctx context.Context) (uint, error) {
	return s.repo.Count(ctx)
}*/

// Query returns the items with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit uint) ([]Product, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of products by ctx")
	}
	return items, nil
}


// List returns the items list.
func (s service) List(ctx context.Context) ([]Product, error) {
	items, err := s.repo.Query(ctx, 0, MaxLIstLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of products by ctx")
	}
	return items, nil
}

func (s service) Create(ctx context.Context, entity *Product) error {
	return s.repo.Create(ctx, entity)
}

func (s service) Update(ctx context.Context, entity *Product) error {
	return s.repo.Update(ctx, entity)
}

func (s service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s service) BuildByConditons(product *Product, conditions []condition.Condition) (*Product, error) {
	p := &Product{
		Name:	product.Name,
	}
	components := make([]component.Component, 0)

	for _, c := range product.Components {
		co, err := s.componentService.BuildByConditons(&c, conditions)
		if err == nil {
			components = append(components, *co)
		}
	}
	p.Components = components

	if !p.IsValid() {
		return nil, apperror.ErrNotFound
	}
	return p, nil
}

