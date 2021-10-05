package service

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
)

type CartService interface {
	Checkout(ctx context.Context, ID string, items []*entity.CartItem) error
	Add(ctx context.Context, items *entity.CartItem) error
}
