package entity_test

import (
	"bytes"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	engine2 "github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/kennykarnama/checkout-challenge/cart/entity"
	stockEntity "github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__RuleCheckoutMacbookRaspberry(t *testing.T) {
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Checkout Calculator", "0.0.1", pkg.NewFileResource("CheckoutRule.grl"))
	assert.NoError(t, err)

	engine := engine2.NewGruleEngine()

	kb := lib.NewKnowledgeBaseInstance("Checkout Calculator", "0.0.1")

	buff := &bytes.Buffer{}
	cat := kb.MakeCatalog()
	err = cat.WriteCatalogToWriter(buff)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff.Bytes())
	cat2 := &ast.Catalog{}
	cat2.ReadCatalogFromReader(buff2)
	nkb := cat2.BuildKnowledgeBase()

	cartItems := []*entity.CartItem{
		{
			ID:  "KENNY",
			SKU: "43N23P",
			Qty: 1,
		},
		{
			ID:  "KENNY",
			SKU: "234234",
			Qty: 1,
		},
	}
	mappedCart := entity.CartItems(cartItems).Normalize()
	mappedSku := make(map[string]*stockEntity.StockItem)
	mappedSku["43N23P"] = &stockEntity.StockItem{
		SKU:          "43N23P",
		Name:         "Macbook",
		Price:        5399.99,
		InventoryQty: 1,
	}
	mappedSku["234234"] = &stockEntity.StockItem{
		SKU:          "234234",
		Name:         "Raspberry PI",
		Price:        30,
		InventoryQty: 1,
	}
	checkout := &entity.Checkout{TotalPrice: 0, CurrencySymbol: "$"}

	for _, cartItem := range cartItems {
		dctx := ast.NewDataContext()
		dctx.Add("CartItem", cartItem)
		dctx.Add("Checkout", checkout)
		dctx.Add("MappedCartItem", mappedCart)
		dctx.Add("MappedSku", mappedSku)
		err = engine.Execute(dctx, nkb)
		assert.NoError(t, err)
	}

	expected := "$5399.99"

	if expected != checkout.String() {
		t.Errorf("expected total: %v but found: %v", expected, checkout.String())
		return
	}
}

func Test__RuleGoogleHome(t *testing.T) {
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Checkout Calculator", "0.0.1", pkg.NewFileResource("CheckoutRule.grl"))
	assert.NoError(t, err)

	engine := engine2.NewGruleEngine()

	kb := lib.NewKnowledgeBaseInstance("Checkout Calculator", "0.0.1")

	buff := &bytes.Buffer{}
	cat := kb.MakeCatalog()
	err = cat.WriteCatalogToWriter(buff)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff.Bytes())
	cat2 := &ast.Catalog{}
	cat2.ReadCatalogFromReader(buff2)
	nkb := cat2.BuildKnowledgeBase()

	cartItems := []*entity.CartItem{
		{
			ID:  "KENNY",
			SKU: "120P90",
			Qty: 3,
		},
	}
	mappedCart := entity.CartItems(cartItems).Normalize()
	mappedSku := make(map[string]*stockEntity.StockItem)
	mappedSku["120P90"] = &stockEntity.StockItem{
		SKU:          "120P90",
		Name:         "Google Home",
		Price:        49.99,
		InventoryQty: 3,
	}
	checkout := &entity.Checkout{TotalPrice: 0, CurrencySymbol: "$"}

	for _, cartItem := range cartItems {
		dctx := ast.NewDataContext()
		dctx.Add("CartItem", cartItem)
		dctx.Add("Checkout", checkout)
		dctx.Add("MappedCartItem", mappedCart)
		dctx.Add("MappedSku", mappedSku)
		err = engine.Execute(dctx, nkb)
		assert.NoError(t, err)
	}

	expected := "$99.98"

	if expected != checkout.String() {
		t.Errorf("expected total: %v but found: %v", expected, checkout.String())
		return
	}
}

func Test__RuleAlexaSpeaker(t *testing.T) {
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Checkout Calculator", "0.0.1", pkg.NewFileResource("CheckoutRule.grl"))
	assert.NoError(t, err)

	engine := engine2.NewGruleEngine()

	kb := lib.NewKnowledgeBaseInstance("Checkout Calculator", "0.0.1")

	buff := &bytes.Buffer{}
	cat := kb.MakeCatalog()
	err = cat.WriteCatalogToWriter(buff)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff.Bytes())
	cat2 := &ast.Catalog{}
	cat2.ReadCatalogFromReader(buff2)
	nkb := cat2.BuildKnowledgeBase()

	cartItems := []*entity.CartItem{
		{
			ID:  "KENNY",
			SKU: "A304SD",
			Qty: 3,
		},
	}
	mappedCart := entity.CartItems(cartItems).Normalize()
	mappedSku := make(map[string]*stockEntity.StockItem)
	mappedSku["A304SD"] = &stockEntity.StockItem{
		SKU:          "A304SD",
		Name:         "Alexa Speaker",
		Price:        109.50,
		InventoryQty: 3,
	}
	checkout := &entity.Checkout{TotalPrice: 0, CurrencySymbol: "$"}

	for _, cartItem := range cartItems {
		dctx := ast.NewDataContext()
		dctx.Add("CartItem", cartItem)
		dctx.Add("Checkout", checkout)
		dctx.Add("MappedCartItem", mappedCart)
		dctx.Add("MappedSku", mappedSku)
		err = engine.Execute(dctx, nkb)
		assert.NoError(t, err)
	}

	expected := "$295.65"

	if expected != checkout.String() {
		t.Errorf("expected total: %v but found: %v", expected, checkout.String())
		return
	}
}

func Test__RuleOtherwise(t *testing.T) {
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	err := rb.BuildRuleFromResource("Checkout Calculator", "0.0.1", pkg.NewFileResource("CheckoutRule.grl"))
	assert.NoError(t, err)

	engine := engine2.NewGruleEngine()

	kb := lib.NewKnowledgeBaseInstance("Checkout Calculator", "0.0.1")

	buff := &bytes.Buffer{}
	cat := kb.MakeCatalog()
	err = cat.WriteCatalogToWriter(buff)
	assert.Nil(t, err)

	buff2 := bytes.NewBuffer(buff.Bytes())
	cat2 := &ast.Catalog{}
	cat2.ReadCatalogFromReader(buff2)
	nkb := cat2.BuildKnowledgeBase()

	cartItems := []*entity.CartItem{
		{
			ID:  "KENNY",
			SKU: "A304SD",
			Qty: 1,
		},
		{
			ID:  "KENNY",
			SKU: "120P90",
			Qty: 1,
		},
	}
	mappedCart := entity.CartItems(cartItems).Normalize()
	mappedSku := make(map[string]*stockEntity.StockItem)
	mappedSku["A304SD"] = &stockEntity.StockItem{
		SKU:          "A304SD",
		Name:         "Alexa Speaker",
		Price:        109.50,
		InventoryQty: 3,
	}
	mappedSku["120P90"] = &stockEntity.StockItem{
		SKU:          "120P90",
		Name:         "Google Speaker",
		Price:        49.99,
		InventoryQty: 3,
	}
	checkout := &entity.Checkout{TotalPrice: 0, CurrencySymbol: "$"}

	for _, cartItem := range cartItems {
		dctx := ast.NewDataContext()
		dctx.Add("CartItem", cartItem)
		dctx.Add("Checkout", checkout)
		dctx.Add("MappedCartItem", mappedCart)
		dctx.Add("MappedSku", mappedSku)
		err = engine.Execute(dctx, nkb)
		assert.NoError(t, err)
	}

	expected := "$159.49"

	if expected != checkout.String() {
		t.Errorf("expected total: %v but found: %v", expected, checkout.String())
		return
	}
}
