package db

import (
	"context"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"
)

// ProductRepository is a repository for the product entity
type RuleApplicabilityRepository struct {
	repository
}

var _ ruleapplicability.IRepository = (*RuleApplicabilityRepository)(nil)

// New creates a new Repository
func NewRuleApplicabilityRepository(repository *repository) (*RuleApplicabilityRepository, error) {
	return &RuleApplicabilityRepository{repository: *repository}, nil
}


// Get reads entities with the specified ID from the database.
func (r RuleApplicabilityRepository) Get(ctx context.Context, id uint) (*ruleapplicability.RuleApplicability, error) {
	entity := &ruleapplicability.RuleApplicability{}

	err := r.dbWithDefaults().First(&entity, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r RuleApplicabilityRepository) First(ctx context.Context, entity *ruleapplicability.RuleApplicability) (*ruleapplicability.RuleApplicability, error) {
	err := r.db.DB().Where(entity).First(entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves records with the specified offset and limit from the database.
func (r RuleApplicabilityRepository) Query(ctx context.Context, offset, limit uint) ([]ruleapplicability.RuleApplicability, error) {
	var items []ruleapplicability.RuleApplicability

	err := r.dbWithContext(ctx, r.dbWithDefaults()).Find(&items).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return items, apperror.ErrNotFound
		}
	}
	return items, err
}

// Create saves a new record in the database.
// It returns the ID of the newly inserted record.
func (r RuleApplicabilityRepository) Create(ctx context.Context, entity *ruleapplicability.RuleApplicability) error {

	if !r.db.DB().NewRecord(entity) {
		return errors.New("entity is not new")
	}
	return r.db.DB().Create(entity).Error
}

func (r RuleApplicabilityRepository) Update(ctx context.Context, entity *ruleapplicability.RuleApplicability) error {

	if r.db.DB().NewRecord(entity) {
		return errors.New("entity is new")
	}
	return r.db.DB().Save(entity).Error
}

// Delete deletes a record with the specified ID from the database.
func (r RuleApplicabilityRepository) Delete(ctx context.Context, id uint) error {
	entity, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.DB().Delete(entity).Error
}

