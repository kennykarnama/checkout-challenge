package mapbased

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/kennykarnama/checkout-challenge/stock/repository"
	"sync"
)

type MemoryMappedRepository struct {
	data sync.Map
	mut  *sync.Mutex
}

func NewMemoryMappedRepository(options ...Option) *MemoryMappedRepository {
	m := &MemoryMappedRepository{
		data: sync.Map{},
		mut:  &sync.Mutex{},
	}
	for _, opt := range options {
		opt(m)
	}
	return m
}

func (m *MemoryMappedRepository) GetStocksBySKU(ctx context.Context, skus []string) ([]*entity.StockItem, error) {
	stockItems := []*entity.StockItem{}
	for _, sku := range skus {
		raw, ok := m.data.Load(sku)
		if ok {
			stockItem, valid := raw.(*entity.StockItem)
			if valid {
				stockItems = append(stockItems, stockItem)
			}
		}
	}
	return stockItems, nil
}

func (m *MemoryMappedRepository) DecrementStockQtyBySku(ctx context.Context, sku string, desiredQty int64) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	item, _ := m.GetStockBySKU(ctx, sku)
	item.InventoryQty -= desiredQty
	if item.InventoryQty < 0 {
		return repository.ErrInventoryQtyNotSufficient
	}
	m.data.Store(sku, item)
	return nil
}

func (m *MemoryMappedRepository) AddStock(ctx context.Context, newItem *entity.StockItem) error {
	m.data.Store(newItem.SKU, newItem)
	return nil
}

func (m *MemoryMappedRepository) GetStockBySKU(ctx context.Context, sku string) (*entity.StockItem, error) {
	raw, ok := m.data.Load(sku)
	if ok {
		stockItem, valid := raw.(*entity.StockItem)
		if valid {
			return stockItem, nil
		}
	}
	return nil, nil
}
