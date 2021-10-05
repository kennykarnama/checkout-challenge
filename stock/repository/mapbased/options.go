package mapbased

import "github.com/kennykarnama/checkout-challenge/stock/entity"

type Option func(*MemoryMappedRepository)

func WithData(items []*entity.StockItem) Option {
	return func(repository *MemoryMappedRepository) {
		for _, item := range items {
			v := item
			repository.data.Store(item.SKU, v)
		}
	}
}
