package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
)

// ProductRepository is a repository for the product entity
type PriceRepository struct {
	repository
}

var _ price.IRepository = (*PriceRepository)(nil)

// New creates a new Repository
func NewPriceRepository(repository *repository) (*PriceRepository, error) {
	return &PriceRepository{repository: *repository}, nil
}


// Get reads entities with the specified ID from the database.
func (r PriceRepository) Get(ctx context.Context, id uint) (*price.Price, error) {
	entity := &price.Price{}

	err := r.dbWithDefaults().First(&entity, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r PriceRepository) First(ctx context.Context, entity *price.Price) (*price.Price, error) {
	err := r.db.DB().Where(entity).First(entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves records with the specified offset and limit from the database.
func (r PriceRepository) Query(ctx context.Context, offset, limit uint) ([]price.Price, error) {
	var items []price.Price

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
func (r PriceRepository) Create(ctx context.Context, entity *price.Price) error {

	if !r.db.DB().NewRecord(entity) {
		return errors.New("entity is not new")
	}
	return r.db.DB().Create(entity).Error
}

func (r PriceRepository) Update(ctx context.Context, entity *price.Price) error {

	if r.db.DB().NewRecord(entity) {
		return errors.New("entity is new")
	}
	return r.db.DB().Save(entity).Error
}

// Delete deletes a record with the specified ID from the database.
func (r PriceRepository) Delete(ctx context.Context, id uint) error {
	entity, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.DB().Delete(entity).Error
}

