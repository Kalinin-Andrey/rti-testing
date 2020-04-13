package ruleapplicability

import (
	"context"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"github.com/pkg/errors"

	"github.com/Kalinin-Andrey/rti-testing/pkg/log"
)

const MaxLIstLimit = 1000

// IService encapsulates usecase logic for RuleApplicability.
type IService interface {
	NewEntity() *RuleApplicability
	Get(ctx context.Context, id uint) (*RuleApplicability, error)
	First(ctx context.Context, entity *RuleApplicability) (*RuleApplicability, error)
	Query(ctx context.Context, offset, limit uint) ([]RuleApplicability, error)
	List(ctx context.Context) ([]RuleApplicability, error)
	//Count(ctx context.Context) (uint, error)
	Create(ctx context.Context, entity *RuleApplicability) error
	Update(ctx context.Context, entity *RuleApplicability) error
	Delete(ctx context.Context, id uint) (error)
	IsSatisfyConditions(r *RuleApplicability, conditions []condition.Condition) (IsSatisfy bool)
}

type service struct {
	//Domain     Domain
	repo       IRepository
	logger     log.ILogger
}

// NewService creates a new service.
func NewService(repo IRepository, logger log.ILogger) IService {
	s := &service{repo, logger}
	repo.SetDefaultConditions(s.defaultConditions())
	return s
}

// Defaults returns defaults params
func (s service) defaultConditions() map[string]interface{} {
	return map[string]interface{}{
	}
}

func (s service) NewEntity() *RuleApplicability {
	return &RuleApplicability{}
}

// Get returns the entity with the specified ID.
func (s service) Get(ctx context.Context, id uint) (*RuleApplicability, error) {
	entity, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not get a RuleApplicability by id: %v", id)
	}
	return entity, nil
}

func (s service) First(ctx context.Context, entity *RuleApplicability) (*RuleApplicability, error) {
	return s.repo.First(ctx, entity)
}

/*
// Count returns the number of items.
func (s service) Count(ctx context.Context) (uint, error) {
	return s.repo.Count(ctx)
}*/

// Query returns the items with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit uint) ([]RuleApplicability, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of RuleApplicabilities by ctx")
	}
	return items, nil
}


// List returns the items list.
func (s service) List(ctx context.Context) ([]RuleApplicability, error) {
	items, err := s.repo.Query(ctx, 0, MaxLIstLimit)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of RuleApplicabilities by ctx")
	}
	return items, nil
}

func (s service) Create(ctx context.Context, entity *RuleApplicability) error {
	return s.repo.Create(ctx, entity)
}

func (s service) Update(ctx context.Context, entity *RuleApplicability) error {
	return s.repo.Update(ctx, entity)
}

func (s service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}


func (s service) IsSatisfyConditions(r *RuleApplicability, conditions []condition.Condition) (IsSatisfy bool) {
	var err error
	for _, c := range conditions {
		IsSatisfy, err =  r.IsSatisfy(c)
		if err != nil {
			s.logger.Error(err)
			continue
		}
		if IsSatisfy {
			IsSatisfy = true
			break
		}
	}

	return IsSatisfy
}

