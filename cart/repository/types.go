package repository

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
)

type CartRepository interface {
	Add(ctx context.Context, item *entity.CartItem) error
	GetCartByID(ctx context.Context, ID string) ([]*entity.CartItem, error)
}
