[![Build Status](https://app.travis-ci.com/kennykarnama/checkout-challenge.svg?branch=main)](https://app.travis-ci.com/kennykarnama/checkout-challenge)

## Project Overview

This is a small project to simulate in simple way the way we add items and checkout from cart.

Limitation

* [ ] Doesn't support custom data loading (we can add it later on)
* [ ] Currency formatting is tied to `$`

This project also assumes input given by user is in list of SKU instead of items' name. 

Rationale:

- SKU give consistent behavior because we don't need to handle spacing or other useless formatting if we use item.name instead
- data consistency can be guaranteed

Also, this project use special identifier called `ID` to group items of cart. If we correlate with real cases, we might found that `ID` can correpond to `userID` since a cart is owned by a user.

## Project Structure

This project consists two main packages

### stock package

This package is responsible for mantaining the item information details

This package has an entity which is structured like this

```go
type StockItem struct {
	SKU          string
	Name         string
	Price        float64
	InventoryQty int64
}
```

Under this package we have repository & service to mantain data flow

### cart package

This package is responsible for handling activities:

- Add item to cart
- Checkout

Also in this package, checkout price was calculated based on `grl` files

This is an example of grl file

```
rule MacbookRule "When you buy Macbook, you got Raspberry PI B For free" {
    when
        CartItem.SKU == "43N23P" && MappedCartItem["234234"].Qty > 0
    then
        MappedSku["234234"].Price = 0;
        Retract("MacbookRule");
}

rule GoogleSpeaker "When you buy 3 google speaker, only pay for 2" {
    when
        CartItem.SKU == "120P90" && MappedCartItem[CartItem.SKU].Qty >= 3
    then
        MappedCartItem[CartItem.SKU].Qty = MappedCartItem[CartItem.SKU].Qty - 1;
        Retract("GoogleSpeaker");
}

rule AlexSpeaker "When you buy more than 3 alexa speakers, got 10% off on each item of this speakers" {
    when
        CartItem.SKU == "A304SD" && MappedCartItem[CartItem.SKU].Qty >= 3
    then
        MappedSku[CartItem.SKU].Price = MappedSku[CartItem.SKU].Price - (MappedSku[CartItem.SKU].Price * 10 / 100);
        Retract("AlexSpeaker");
}


rule GeneralPrice "Otherwise" {
    when
        MappedCartItem[CartItem.SKU].Qty > 0
    then
        Log(MappedSku[CartItem.SKU].String());
        Checkout.TotalPrice = Checkout.TotalPrice + MappedSku[CartItem.SKU].Price * CartItem.Qty;
        MappedCartItem[CartItem.SKU].Qty = 0;
}

```

This `.grl` file will act as knowledge base to determine what price should be paid by the customer.

## Unit Test

Under each of the package, i add some test cases. You can find in either of the following files or folders

- `test/`
- `[file]_test.go`

## Run the project

This project doesn't support sophisticated data mantaining. It wants to as simple as it could be.

So to run this project, you should prepare a file. The file has the following structure

```
FIRST LINE will be ID
[SKU] 2 .. N LINE will be SKU
```

example

```
KENNY
43N23P
234234
```

To process it, simply run this command on your terminal

```
go run . -input .\in.txt
```

It will print the total checkout price

```
Total: $5399.99
```