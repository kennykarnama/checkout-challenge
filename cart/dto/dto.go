package dto

import "github.com/kennykarnama/checkout-challenge/cart/entity"

type AddToCartRequet struct {
	Items []*entity.CartItem
}

type AddToChartResponse struct {
	ID string
}
