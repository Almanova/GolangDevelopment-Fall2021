package inmemory

import (
	"context"
	"fmt"
	"project/internal/models"
	"project/internal/store"
	"sync"
)

type ProductsRepo struct {
	data map[int]*models.ProductDto
	categoriesRepo store.CategoryRepository
	brandsRepo store.BrandRepository

	mu *sync.RWMutex
}


func (db *ProductsRepo) Create(ctx context.Context, product *models.ProductDto) (*models.ProductDto, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	product.Id = len(db.data) + 1
	db.data[product.Id] = product
	return product, nil
}

func (db *ProductsRepo) All(ctx context.Context, categoryId int, brandId int) ([]*models.Product, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	products := make([]*models.Product, 0, len(db.data))
	for _, product := range db.data {
		category, err := db.categoriesRepo.ByID(ctx, product.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("internal server error occured")
		}
		brand, err := db.brandsRepo.ByID(ctx, product.BrandId)
		if err != nil {
			return nil, fmt.Errorf("internal server error occured")
		}
		if (product.CategoryId == categoryId || categoryId == 0) &&
			(product.BrandId == brandId || brandId == 0) {
			products = append(products, &models.Product{
				Id: product.Id,
				Name: product.Name,
				CategoryName: category.Name,
				BrandName: brand.Name,
				Description: product.Description,
				Weight: product.Weight,
				Price: product.Price,
				Discount: product.Discount,
			})
		}
	}

	return products, nil
}

func (db *ProductsRepo) ByID(ctx context.Context, id int) (*models.ProductFull, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	product, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("no product with id %d", id)
	}

	category, err := db.categoriesRepo.ByID(ctx, product.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("internal server error occured")
	}
	brand, err := db.brandsRepo.ByID(ctx, product.BrandId)
	if err != nil {
		return nil, fmt.Errorf("internal server error occured")
	}

	return &models.ProductFull{
		Id: product.Id,
		Name: product.Name,
		Category: *category,
		Brand: *brand,
		Description: product.Description,
		Weight: product.Weight,
		Price: product.Price,
		Discount: product.Discount,
	}, nil
}

func (db *ProductsRepo) Update(ctx context.Context, product *models.ProductDto, id int) (*models.ProductDto, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	product.Id = id
	db.data[id] = product
	return product, nil
}

func (db *ProductsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
