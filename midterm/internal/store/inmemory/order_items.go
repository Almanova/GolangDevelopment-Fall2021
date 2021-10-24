package inmemory

import (
	"context"
	"project/internal/models"
	"project/internal/store"
	"sync"
)

type OrderItemsRepo struct {
	data map[int]*models.OrderItem
	productsRepo store.ProductRepository

	mu *sync.RWMutex
}

func (db *OrderItemsRepo) Create(ctx context.Context, cartItem *models.CartItemFull, orderId int) (*models.OrderItem, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	orderItem := &models.OrderItem{
		Id: len(db.data) + 1,
		OrderId: orderId,
		ProductId: cartItem.Product.Id,
		Quantity: cartItem.Quantity,
		Price: cartItem.Product.Price,
		Discount: cartItem.Product.Discount,
		TotalPrice: (float64(cartItem.Product.Price) - float64(cartItem.Product.Price) / 100 *
			float64(cartItem.Product.Discount)) * float64(cartItem.Quantity),
	}

	db.data[orderItem.Id] = orderItem
	return orderItem, nil
}
