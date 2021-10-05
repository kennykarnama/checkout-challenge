package service

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/stock/entity"
)

type StockService interface {
	DecrementStockQtyBySku(ctx context.Context, sku string, decrement int64) error
	GetStocksBySKU(ctx context.Context, skus []string) ([]*entity.StockItem, error)
}
