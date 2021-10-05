package entity

import "fmt"

type StockItem struct {
	SKU          string
	Name         string
	Price        float64
	InventoryQty int64
}

type StockItems []*StockItem

func (si StockItems) MappedBySKU() map[string]*StockItem {
	m := make(map[string]*StockItem)
	for _, stockItem := range si {
		v := stockItem
		m[v.SKU] = v
	}
	return m
}

func (s *StockItem) String() string {
	return fmt.Sprintf("%s:%v", s.SKU, s.Price)
}
