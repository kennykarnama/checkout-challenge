package service

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/kennykarnama/checkout-challenge/stock/repository"
)

type StockServiceImpl struct {
	repo repository.StockRepository
}

func NewStockServiceImpl(repo repository.StockRepository) *StockServiceImpl {
	return &StockServiceImpl{repo: repo}
}

func (s *StockServiceImpl) DecrementStockQtyBySku(ctx context.Context, sku string, decrement int64) error {
	return s.repo.DecrementStockQtyBySku(ctx, sku, decrement)
}

func (s *StockServiceImpl) GetStocksBySKU(ctx context.Context, skus []string) ([]*entity.StockItem, error) {
	return s.repo.GetStocksBySKU(ctx, skus)
}
