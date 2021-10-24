package inmemory

import (
	"project/internal/models"
	"project/internal/store"
	"sync"
)

type DB struct {
	productsRepo store.ProductRepository
	categoriesRepo store.CategoryRepository
	brandsRepo store.BrandRepository
	cartItemsRepo store.CartItemRepository
	orderItemsRepo store.OrderItemRepository
	ordersRepo store.OrderRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Products() store.ProductRepository {
	if db.productsRepo == nil {
		db.productsRepo = &ProductsRepo{
			data: make(map[int]*models.ProductDto),
			categoriesRepo: db.Categories(),
			brandsRepo: db.Brands(),
			mu: new(sync.RWMutex),
		}
	}

	return db.productsRepo
}

func (db *DB) Categories() store.CategoryRepository {
	if db.categoriesRepo == nil {
		db.categoriesRepo = &CategoriesRepo{
			data: make(map[int]*models.Category),
			mu: new(sync.RWMutex),
		}
	}

	return db.categoriesRepo
}

func (db *DB) Brands() store.BrandRepository {
	if db.brandsRepo == nil {
		db.brandsRepo = &BrandsRepo{
			data: make(map[int]*models.Brand),
			mu: new(sync.RWMutex),
		}
	}

	return db.brandsRepo
}

func (db *DB) CartItems() store.CartItemRepository {
	if db.cartItemsRepo == nil {
		db.cartItemsRepo = &CartItemsRepo{
			data: make(map[int]*models.CartItem),
			productsRepo: db.Products(),
			mu: new(sync.RWMutex),
		}
	}

	return db.cartItemsRepo
}

func (db *DB) OrderItems() store.OrderItemRepository {
	if db.orderItemsRepo == nil {
		db.orderItemsRepo = &OrderItemsRepo{
			data: make(map[int]*models.OrderItem),
			productsRepo: db.Products(),
			mu: new(sync.RWMutex),
		}
	}

	return db.orderItemsRepo
}

func (db *DB) Orders() store.OrderRepository {
	if db.ordersRepo == nil {
		db.ordersRepo = &OrdersRepo{
			data: make(map[int]*models.Order),
			cartItemsRepo: db.CartItems(),
			orderItemsRepo: db.OrderItems(),
			productsRepo: db.Products(),
			mu: new(sync.RWMutex),
		}
	}

	return db.ordersRepo
}
