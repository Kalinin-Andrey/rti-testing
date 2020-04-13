package price

import (
	"context"
)

// IRepository encapsulates the logic to access enties from the data source.
type IRepository interface {
	// Get returns an entity with the specified ID.
	Get(ctx context.Context, id uint) (*Price, error)
	First(ctx context.Context, user *Price) (*Price, error)
	// Count returns the number of entities.
	//Count(ctx context.Context) (uint, error)
	// Query returns the list of entities with the given offset and limit.
	Query(ctx context.Context, offset, limit uint) ([]Price, error)
	SetDefaultConditions(conditions map[string]interface{})
	// Create saves a new entity in the storage.
	Create(ctx context.Context, entity *Price) error
	// Update updates an entity with given ID in the storage.
	Update(ctx context.Context, entity *Price) error
	// Delete removes an entity with given ID from the storage.
	Delete(ctx context.Context, id uint) error
}

