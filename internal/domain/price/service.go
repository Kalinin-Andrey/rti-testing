package price

import (
	"context"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"
	"github.com/pkg/errors"

	"github.com/Kalinin-Andrey/rti-testing/pkg/log"
)

const MaxLIstLimit = 1000

// IService encapsulates usecase logic for price.
type IService interface {
	NewEntity() *Price
	Get(ctx context.Context, id uint) (*Price, error)
	First(ctx context.Context, entity *Price) (*Price, error)
	Query(ctx context.Context, offset, limit uint) ([]Price, error)
	List(ctx context.Context) ([]Price, error)
	//Count(ctx context.Context) (uint, error)
	Create(ctx context.Context, entity *Price) error
	Update(ctx context.Context, entity *Price) error
	Delete(ctx context.Context, id uint) (error)
	IsSatisfyConditions(price *Price, conditions []condition.Condition) (isSatisfy bool)
}

type service struct {
	//Domain     Domain
	repo       IRepository
	ruleService	ruleapplicability.IService
	logger     log.ILogger
}

// NewService creates a new service.
func NewService(repo IRepository, ruleRepo ruleapplicability.IService, logger log.ILogger) IService {
	s := &service{repo, ruleRepo, logger}
	repo.SetDefaultConditions(s.defaultConditions())
	return s
}

// Defaults returns defaults params
func (s service) defaultConditions() map[string]interface{} {
	return map[string]interface{}{
	}
}

func (s service) NewEntity() *Price {
	return &Price{}
}

// Get returns the entity with the specified ID.
func (s service) Get(ctx context.Context, id uint) (*Price, error) {
	entity, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not get a price by id: %v", id)
	}
	return entity, nil
}
/*
// Count returns the number of items.
func (s service) Count(ctx context.Context) (uint, error) {
	return s.repo.Count(ctx)
}*/

// Query returns the items with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit uint) ([]Price, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of prices by ctx")
	}
	return items, nil
}


// List returns the items list.
func (s service) List(ctx context.Context) ([]Price, error) {
	items, err := s.repo.Query(ctx, 0, MaxLIstLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of prices by ctx")
	}
	return items, nil
}

func (s service) Create(ctx context.Context, entity *Price) error {
	return s.repo.Create(ctx, entity)
}

func (s service) Update(ctx context.Context, entity *Price) error {
	return s.repo.Update(ctx, entity)
}

func (s service) First(ctx context.Context, entity *Price) (*Price, error) {
	return s.repo.First(ctx, entity)
}

func (s service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s service) IsSatisfyConditions(price *Price, conditions []condition.Condition) (isSatisfy bool) {

	for _, r := range price.RuleApplicabilities {
		if !s.ruleService.IsSatisfyConditions(&r, conditions) {
			return false
		}
	}

	if !price.IsValid() {
		return false
	}
	return true
}
