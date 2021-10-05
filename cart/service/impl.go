package service

import (
	"bytes"
	"context"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	engine2 "github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
	"github.com/kennykarnama/checkout-challenge/cart/repository"
	stockEntity "github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/kennykarnama/checkout-challenge/stock/service"
)

type CartServiceImpl struct {
	stockService     service.StockService
	repo             repository.CartRepository
	CheckoutRuleFile string
}

func NewCartServiceImpl(stockService service.StockService, repo repository.CartRepository, checkoutRuleFile string) *CartServiceImpl {
	return &CartServiceImpl{stockService: stockService, repo: repo, CheckoutRuleFile: checkoutRuleFile}
}

func (s *CartServiceImpl) Checkout(ctx context.Context, ID string) (*entity.Checkout, error) {
	cartItems, err := s.repo.GetCartByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	skuDetails, err := s.stockService.GetStocksBySKU(ctx, entity.CartItems(cartItems).SKUs())
	if err != nil {
		return nil, err
	}

	mappedSkuDetails := stockEntity.StockItems(skuDetails).MappedBySKU()

	mappedCartItems := entity.CartItems(cartItems).Normalize()

	checkout := &entity.Checkout{
		CurrencySymbol: "$",
	}

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err = rb.BuildRuleFromResource("Checkout Calculator", "0.0.1", pkg.NewFileResource("CheckoutRule.grl"))
	if err != nil {
		return nil, err
	}

	engine := engine2.NewGruleEngine()

	kb := lib.NewKnowledgeBaseInstance("Checkout Calculator", "0.0.1")

	buff := &bytes.Buffer{}
	cat := kb.MakeCatalog()
	err = cat.WriteCatalogToWriter(buff)
	if err != nil {
		return nil, err
	}

	buff2 := bytes.NewBuffer(buff.Bytes())
	cat2 := &ast.Catalog{}
	cat2.ReadCatalogFromReader(buff2)
	nkb := cat2.BuildKnowledgeBase()

	for _, cartItem := range cartItems {
		dctx := ast.NewDataContext()
		dctx.Add("CartItem", cartItem)
		dctx.Add("Checkout", checkout)
		dctx.Add("MappedCartItem", mappedCartItems)
		dctx.Add("MappedSku", mappedSkuDetails)
		err = engine.Execute(dctx, nkb)
	}

	return checkout, nil
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
