package service

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
)

type CartService interface {
	Checkout(ctx context.Context, ID string) (*entity.Checkout, error)
	Add(ctx context.Context, items *entity.CartItem) error
}
