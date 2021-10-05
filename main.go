package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	cartEntity "github.com/kennykarnama/checkout-challenge/cart/entity"
	cartMapBased "github.com/kennykarnama/checkout-challenge/cart/repository/mapbased"
	cartService "github.com/kennykarnama/checkout-challenge/cart/service"
	"github.com/kennykarnama/checkout-challenge/stock/entity"
	"github.com/kennykarnama/checkout-challenge/stock/repository/mapbased"
	"github.com/kennykarnama/checkout-challenge/stock/service"
	"log"
	"os"
)

func main() {
	stockRepository := mapbased.NewMemoryMappedRepository(mapbased.WithData([]*entity.StockItem{
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
	stockService := service.NewStockServiceImpl(stockRepository)

	cartRepo := cartMapBased.NewMemoryMappedRepository()
	cartService := cartService.NewCartServiceImpl(stockService, cartRepo, "CheckoutRule.grl")

	inputFile := flag.String("input", "in.txt", "input file, default is in.txt")
	flag.Parse()

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	/**
	File structure
	line 1 --> ID (can be anything, but in our case, i assume will be userID)
	line 2..n --> SKU of items ordered
	*/
	userID := ""
	lineNum := 1
	for scanner.Scan() {
		if lineNum == 1 {
			userID = scanner.Text()
		} else {
			err = cartService.Add(context.Background(), &cartEntity.CartItem{
				ID:  userID,
				SKU: scanner.Text(),
				Qty: 1,
			})
			if err != nil {
				panic(err)
			}
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// get the result

	checkout, err := cartService.Checkout(context.Background(), userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total: %v\n", checkout.String())
}
