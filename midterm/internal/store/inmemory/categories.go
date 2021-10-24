package inmemory

import (
	"context"
	"fmt"
	"project/internal/models"
	"sync"
)

type CategoriesRepo struct {
	data map[int]*models.Category

	mu *sync.RWMutex
}

func (db *CategoriesRepo) Create(ctx context.Context, category *models.Category) (*models.Category, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	category.Id = len(db.data) + 1
	db.data[category.Id] = category
	return category, nil
}

func (db *CategoriesRepo) All(ctx context.Context) ([]*models.Category, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	categories := make([]*models.Category, 0, len(db.data))
	for _, category := range db.data {
		categories = append(categories, category)
	}

	return categories, nil
}

func (db *CategoriesRepo) ByID(ctx context.Context, id int) (*models.Category, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	category, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no category with id %d", id)
	}

	return category, nil
}

func (db *CategoriesRepo) Update(ctx context.Context, category *models.Category, id int) (*models.Category, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	category.Id = id
	db.data[id] = category
	return category, nil
}

func (db *CategoriesRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
