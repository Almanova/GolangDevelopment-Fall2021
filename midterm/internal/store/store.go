package store

import (
	"context"
	"project/internal/models"
)

type Store interface {
	Products() ProductRepository
	Categories() CategoryRepository
	Brands() BrandRepository
	CartItems() CartItemRepository
	OrderItems() OrderItemRepository
	Orders() OrderRepository
}

type ProductRepository interface {
	Create(ctx context.Context, product *models.ProductDto) (*models.ProductDto, error)
	All(ctx context.Context, categoryId int, brandId int) ([]*models.Product, error)
	ByID(ctx context.Context, id int) (*models.ProductFull, error)
	Update(ctx context.Context, product *models.ProductDto, id int) (*models.ProductDto, error)
	Delete(ctx context.Context, id int) error
}

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) (*models.Category, error)
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category, id int) (*models.Category, error)
	Delete(ctx context.Context, id int) error
}

type BrandRepository interface {
	Create(ctx context.Context, brand *models.Brand) (*models.Brand, error)
	All(ctx context.Context) ([]*models.Brand, error)
	ByID(ctx context.Context, id int) (*models.Brand, error)
	Update(ctx context.Context, brand *models.Brand, id int) (*models.Brand, error)
	Delete(ctx context.Context, id int) error
}

type CartItemRepository interface {
	Create(ctx context.Context, brand *models.CartItem) (*models.CartItem, error)
	ListByUserId(ctx context.Context, userId int) ([]*models.CartItemFull, error)
	Update(ctx context.Context, brand *models.CartItem, id int) (*models.CartItem, error)
	Delete(ctx context.Context, id int) error
}

type OrderItemRepository interface {
	Create(ctx context.Context, cartItem *models.CartItemFull, orderId int) (*models.OrderItem, error)
}

type OrderRepository interface {
	Create(ctx context.Context, orderItem *models.OrderDto) (*models.Order, error)
	ListByUserId(ctx context.Context, userId int) ([]*models.Order, error)
}
