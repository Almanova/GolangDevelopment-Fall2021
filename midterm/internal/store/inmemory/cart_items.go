package inmemory

import (
	"context"
	"fmt"
	"project/internal/models"
	"project/internal/store"
	"sync"
)

type CartItemsRepo struct {
	data map[int]*models.CartItem
	productsRepo store.ProductRepository

	mu *sync.RWMutex
}

func (db *CartItemsRepo) Create(ctx context.Context, cartItem *models.CartItem) (*models.CartItem, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	cartItem.Id = len(db.data) + 1
	db.data[cartItem.Id] = cartItem
	return cartItem, nil
}

func (db *CartItemsRepo) ListByUserId(ctx context.Context, userId int) ([]*models.CartItemFull, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	cartItems := make([]*models.CartItemFull, 0, len(db.data))
	for _, cartItem := range db.data {
		product, err := db.productsRepo.ByID(ctx, cartItem.ProductId)
		if err != nil {
			return nil, fmt.Errorf("internal server error occured")
		}
		if cartItem.UserId == userId {
			cartItems = append(cartItems, &models.CartItemFull{
				Id: cartItem.Id,
				UserId: cartItem.UserId,
				Product: *product,
				Quantity: cartItem.Quantity,
			})
		}
	}

	return cartItems, nil
}

func (db *CartItemsRepo) Update(ctx context.Context, cartItem *models.CartItem, id int) (*models.CartItem, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	cartItem.Id = id
	db.data[id] = cartItem
	return cartItem, nil
}

func (db *CartItemsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
