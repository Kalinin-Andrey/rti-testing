package db

import (
	"context"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"
)

// ProductRepository is a repository for the product entity
type ProductRepository struct {
	repository
}

var _ product.IRepository = (*ProductRepository)(nil)

// New creates a new Repository
func NewProductRepository(repository *repository) (*ProductRepository, error) {
	return &ProductRepository{repository: *repository}, nil
}


// Get reads entities with the specified ID from the database.
func (r ProductRepository) Get(ctx context.Context, id uint) (*product.Product, error) {
	entity := &product.Product{}

	err := r.dbWithDefaults().First(&entity, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r ProductRepository) First(ctx context.Context, entity *product.Product) (*product.Product, error) {
	err := r.db.DB().Where(entity).First(entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves records with the specified offset and limit from the database.
func (r ProductRepository) Query(ctx context.Context, offset, limit uint) ([]product.Product, error) {
	var items []product.Product

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
func (r ProductRepository) Create(ctx context.Context, entity *product.Product) error {

	if !r.db.DB().NewRecord(entity) {
		return errors.New("entity is not new")
	}
	return r.db.DB().Create(entity).Error
}

func (r ProductRepository) Update(ctx context.Context, entity *product.Product) error {

	if r.db.DB().NewRecord(entity) {
		return errors.New("entity is new")
	}
	return r.db.DB().Save(entity).Error
}

// Delete deletes a record with the specified ID from the database.
func (r ProductRepository) Delete(ctx context.Context, id uint) error {
	entity, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.DB().Delete(entity).Error
}

