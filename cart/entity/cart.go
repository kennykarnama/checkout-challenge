package entity

import (
	"fmt"
)

type CartItem struct {
	ID  string
	SKU string
	Qty int64
}

type CartItems []*CartItem

func (ci CartItems) SKUs() []string {
	skus := []string{}
	for _, cartItem := range ci {
		skus = append(skus, cartItem.SKU)
	}
	return skus
}

func (i *CartItem) String() string {
	return fmt.Sprintf("item{SKU: %s ID: %v}", i.SKU, i.ID)
}

type MappedCartItem map[string]*CartItem

func (ci CartItems) Normalize() map[string]*CartItem {
	mapped := make(map[string]*CartItem)

	for _, cartItem := range ci {
		v := cartItem
		existing, ok := mapped[cartItem.SKU]
		if !ok {
			mapped[cartItem.SKU] = v
		} else {
			existing.Qty += v.Qty
			mapped[cartItem.SKU] = existing
		}
	}
	return mapped
}

func (m MappedCartItem) Get(sku string) *CartItem {
	val, ok := m[sku]
	if !ok {
		return &CartItem{}
	}
	return val
}
