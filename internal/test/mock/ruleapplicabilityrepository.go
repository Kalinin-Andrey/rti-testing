package mock

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"
)

// UserRepository is a mock for UserRepository
type RuleApplicabilityRepository struct {
	Response struct {
		Get		struct {
			Entity	*ruleapplicability.RuleApplicability
			Err		error
		}
		First	struct {
			Entity	*ruleapplicability.RuleApplicability
			Err		error
		}
		Query	struct {
			List	[]ruleapplicability.RuleApplicability
			Err		error
		}
		Create	struct {
			Entity	*ruleapplicability.RuleApplicability
			Err		error
		}
		Update	struct {
			Entity	*ruleapplicability.RuleApplicability
			Err		error
		}
		Delete	struct {
			Err		error
		}
	}
}

var _ ruleapplicability.IRepository = (*RuleApplicabilityRepository)(nil)

func (r RuleApplicabilityRepository) SetDefaultConditions(conditions map[string]interface{}) {}

func (r RuleApplicabilityRepository) Get(ctx context.Context, id uint) (*ruleapplicability.RuleApplicability, error) {
	return r.Response.Get.Entity, r.Response.Get.Err
}

func (r RuleApplicabilityRepository) First(ctx context.Context, user *ruleapplicability.RuleApplicability) (*ruleapplicability.RuleApplicability, error) {
	return r.Response.First.Entity, r.Response.First.Err
}

func (r RuleApplicabilityRepository) Query(ctx context.Context, offset, limit uint) ([]ruleapplicability.RuleApplicability, error) {
	return r.Response.Query.List, r.Response.Query.Err
}

func (r RuleApplicabilityRepository) Create(ctx context.Context, entity *ruleapplicability.RuleApplicability) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Create.Err
}

func (r RuleApplicabilityRepository) Update(ctx context.Context, entity *ruleapplicability.RuleApplicability) error {
	if r.Response.Create.Entity != nil {
		copier.Copy(&entity, &r.Response.Create.Entity)
	}
	return r.Response.Update.Err
}

func (r RuleApplicabilityRepository) Delete(ctx context.Context, id uint) error {
	return r.Response.Delete.Err
}
