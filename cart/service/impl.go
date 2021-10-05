package service

import (
	"context"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
	"github.com/kennykarnama/checkout-challenge/cart/repository"
	stockEntity "github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/kennykarnama/checkout-challenge/stock/service"
)

type CartServiceImpl struct {
	stockService service.StockService
	repo         repository.CartRepository
}

func NewCartServiceImpl(stockService service.StockService) *CartServiceImpl {
	return &CartServiceImpl{stockService: stockService}
}

func (s *CartServiceImpl) Checkout(ctx context.Context, ID string, items []*entity.CartItem) (float64, error) {
	cartItems, err := s.repo.GetCartByID(ctx, ID)
	if err != nil {
		return 0, err
	}

	skuDetails, err := s.stockService.GetStocksBySKU(ctx, entity.CartItems(cartItems).SKUs())
	if err != nil {
		return 0, err
	}

	mappedSkuDetails := stockEntity.StockItems(skuDetails).MappedBySKU()

	var total float64
	for _, cartItem := range cartItems {
		if val, ok := mappedSkuDetails[cartItem.SKU]; ok {
			total += val.Price * float64(cartItem.Qty)
		}
	}

	return total, nil
}
func (s *CartServiceImpl) Add(ctx context.Context, item *entity.CartItem) error {
	err := s.stockService.DecrementStockQtyBySku(ctx, item.SKU, item.Qty)
	if err != nil {
		return err
	}
	err = s.repo.Add(ctx, item)
	if err != nil {
		return err
	}
	return nil
}
