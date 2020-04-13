package db

import (
	"context"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"
)

// ProductRepository is a repository for the product entity
type ComponentRepository struct {
	repository
}

var _ component.IRepository = (*ComponentRepository)(nil)

// New creates a new Repository
func NewComponentRepository(repository *repository) (*ComponentRepository, error) {
	return &ComponentRepository{repository: *repository}, nil
}


// Get reads entities with the specified ID from the database.
func (r ComponentRepository) Get(ctx context.Context, id uint) (*component.Component, error) {
	entity := &component.Component{}

	err := r.dbWithDefaults().First(&entity, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r ComponentRepository) First(ctx context.Context, entity *component.Component) (*component.Component, error) {
	err := r.db.DB().Where(entity).First(entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves records with the specified offset and limit from the database.
func (r ComponentRepository) Query(ctx context.Context, offset, limit uint) ([]component.Component, error) {
	var items []component.Component

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
func (r ComponentRepository) Create(ctx context.Context, entity *component.Component) error {

	if !r.db.DB().NewRecord(entity) {
		return errors.New("entity is not new")
	}
	return r.db.DB().Create(entity).Error
}

func (r ComponentRepository) Update(ctx context.Context, entity *component.Component) error {

	if r.db.DB().NewRecord(entity) {
		return errors.New("entity is new")
	}
	return r.db.DB().Save(entity).Error
}

// Delete deletes a record with the specified ID from the database.
func (r ComponentRepository) Delete(ctx context.Context, id uint) error {
	entity, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.DB().Delete(entity).Error
}

