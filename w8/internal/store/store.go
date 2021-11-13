package store

import (
	"context"
	"homeworks/w8/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Products() ProductsRepository
}

type ProductsRepository interface {
	Create(ctx context.Context, product *models.Product) error
	All(ctx context.Context) ([]*models.Product, error)
	ByID(ctx context.Context, id int) (*models.Product, error)
	Update(ctx context.Context, product *models.Product, id int) error
	Delete(ctx context.Context, id int) error
}
