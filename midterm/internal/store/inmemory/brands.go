package inmemory

import (
	"context"
	"fmt"
	"project/internal/models"
	"sync"
)

type BrandsRepo struct {
	data map[int]*models.Brand

	mu *sync.RWMutex
}

func (db *BrandsRepo) Create(ctx context.Context, brand *models.Brand) (*models.Brand, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	brand.Id = len(db.data) + 1
	db.data[brand.Id] = brand
	return brand, nil
}

func (db *BrandsRepo) All(ctx context.Context) ([]*models.Brand, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	brands := make([]*models.Brand, 0, len(db.data))
	for _, brand := range db.data {
		brands = append(brands, brand)
	}

	return brands, nil
}

func (db *BrandsRepo) ByID(ctx context.Context, id int) (*models.Brand, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	brand, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no category with id %d", id)
	}

	return brand, nil
}

func (db *BrandsRepo) Update(ctx context.Context, brand *models.Brand, id int) (*models.Brand, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	brand.Id = id
	db.data[id] = brand
	return brand, nil
}

func (db *BrandsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
