package repository

import (
	"context"
	"errors"
	"github.com/kennykarnama/checkout-challenge/stock/entity"
)

var (
	ErrInventoryQtyNotSufficient = errors.New("inventory qty is not sufficient")
)

type StockRepository interface {
	GetStocksBySKU(ctx context.Context, skus []string) ([]*entity.StockItem, error)
	DecrementStockQtyBySku(ctx context.Context, sku string, decrement int64) error
	AddStock(ctx context.Context, newItem *entity.StockItem) error
	GetStockBySKU(ctx context.Context, sku string) (*entity.StockItem, error)
}
