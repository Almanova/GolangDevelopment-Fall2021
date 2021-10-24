package inmemory

import (
	"context"
	"fmt"
	"project/internal/models"
	"project/internal/store"
	"sync"
)

type OrdersRepo struct {
	data map[int]*models.Order
	cartItemsRepo store.CartItemRepository
	orderItemsRepo store.OrderItemRepository
	productsRepo store.ProductRepository

	mu *sync.RWMutex
}

func (db *OrdersRepo) Create(ctx context.Context, order *models.OrderDto) (*models.Order, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	order.Id = len(db.data) + 1
	order.Status = "Confirmed"
	cartItems, err := db.cartItemsRepo.ListByUserId(ctx, order.UserId)
	if err != nil {
		return nil, fmt.Errorf("internal server error occured")
	}
	totalPrice := float64(0)
	products := make([]*models.ProductFull, 0)
	for _, cartItem := range cartItems {
		orderItem, err := db.orderItemsRepo.Create(ctx, cartItem, order.Id)
		if err != nil {
			return nil, fmt.Errorf("internal server error occured")
		}
		product, err := db.productsRepo.ByID(ctx, orderItem.ProductId)
		if err != nil {
			return nil, fmt.Errorf("internal server error occured")
		}
		products = append(products, product)
		db.cartItemsRepo.Delete(ctx, cartItem.Id)
		totalPrice += orderItem.TotalPrice
	}
	order.TotalPrice = totalPrice

	orderFull := &models.Order{
		Id: order.Id,
		UserId: order.UserId,
		Products: products,
		Status: order.Status,
		Address: order.Address,
		TotalPrice: order.TotalPrice,
	}
	db.data[order.Id] = orderFull

	return orderFull, nil
}

func (db *OrdersRepo) ListByUserId(ctx context.Context, userId int) ([]*models.Order, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	orders := make([]*models.Order, 0, len(db.data))
	for _, order := range db.data {
		if order.UserId == userId {
			orders = append(orders,order)
		}
	}

	return orders, nil
}


