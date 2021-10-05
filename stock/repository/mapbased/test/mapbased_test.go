package mapbased__test

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/kennykarnama/checkout-challenge/stock/repository/mapbased"
	"sync"
	"testing"
)

var (
	repo = mapbased.NewMemoryMappedRepository(mapbased.WithData([]*entity.StockItem{
		{
			SKU:          "120P90",
			Name:         "Google Home",
			Price:        49.99,
			InventoryQty: 10,
		},
		{
			SKU:          "43N23P",
			Name:         "MacBook Pro",
			Price:        5399.99,
			InventoryQty: 5,
		},
		{
			SKU:          "A304SD",
			Name:         "Alexa Speaker",
			Price:        109.50,
			InventoryQty: 10,
		},
		{
			SKU:          "234234",
			Name:         "Raspberry Pi B",
			Price:        30,
			InventoryQty: 2,
		},
	}))

	ctx = context.Background()
)

func Test__GetStocksBySKUItem(t *testing.T) {
	desiredSkus := []string{
		"43N23P",
		"234234",
	}
	expectedLen := len(desiredSkus)
	skuItems, _ := repo.GetStocksBySKU(ctx, desiredSkus)
	if len(skuItems) != expectedLen {
		t.Errorf("expected retrieved items to be: %v but found: %v", expectedLen, len(skuItems))
		return
	}
}

func Test__GetStocksBySKUItemMiss(t *testing.T) {
	desiredSkus := []string{
		"43N23P",
		"23423455",
	}
	expectedLen := 1
	skuItems, _ := repo.GetStocksBySKU(ctx, desiredSkus)
	if len(skuItems) != expectedLen {
		t.Errorf("expected retrieved items to be: %v but found: %v", expectedLen, len(skuItems))
		return
	}
}

func Test__GetStocksBySKUItemEmptyQuery(t *testing.T) {
	desiredSkus := []string{}
	expectedLen := len(desiredSkus)
	skuItems, _ := repo.GetStocksBySKU(ctx, desiredSkus)
	if len(skuItems) != expectedLen {
		t.Errorf("expected retrieved items to be: %v but found: %v", expectedLen, len(skuItems))
		return
	}
}

func Test__AddStock(t *testing.T) {
	newItem := &entity.StockItem{
		SKU:          "TEST-SIMPAN",
		Name:         "My Item",
		Price:        12.99,
		InventoryQty: 11,
	}
	repo.AddStock(ctx, newItem)
	retrieved, _ := repo.GetStocksBySKU(ctx, []string{newItem.SKU})
	if len(retrieved) == 0 {
		t.Errorf("expected retrieved item to equal 1")
		return
	}
	item := retrieved[0]
	if item == nil {
		t.Errorf("nil record")
		return
	}
	if !cmp.Equal(*newItem, *item) {
		t.Errorf("found diff between target and retrieved item. diff: %v", cmp.Diff(*newItem, *item))
		return
	}
}

func Test__DecrementStockQtyBySKU(t *testing.T) {
	type updateData struct {
		SKU       string
		Decrement int64
	}

	updateDatas := []updateData{
		{
			SKU:       "43N23P",
			Decrement: 1,
		},
		{
			SKU:       "43N23P",
			Decrement: 1,
		},
		{
			SKU:       "43N23P",
			Decrement: 1,
		},
	}

	var wg sync.WaitGroup

	totalDiff := make(map[string]int64)
	expectedResult := make(map[string]int64)

	for _, ud := range updateDatas {
		totalDiff[ud.SKU] += ud.Decrement
	}

	for _, ud := range updateDatas {
		item, _ := repo.GetStockBySKU(ctx, ud.SKU)
		expectedResult[ud.SKU] = item.InventoryQty - totalDiff[ud.SKU]
	}

	for _, ud := range updateDatas {
		totalDiff[ud.SKU] += ud.Decrement

		wg.Add(1)

		go func(target updateData) {
			repo.DecrementStockQtyBySku(ctx, target.SKU, target.Decrement)
			defer wg.Done()
		}(ud)

	}

	wg.Wait()

	for sku, expected := range expectedResult {
		item, _ := repo.GetStockBySKU(ctx, sku)
		if item == nil {
			t.Errorf("nil record on sku: %v", sku)
			return
		}
		if item.InventoryQty != expected {
			t.Errorf("failed on sku: %v. expected inventory qty to be: %v but found: %v", sku, expected, item.InventoryQty)
			return
		}

	}
}
