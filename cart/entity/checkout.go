package entity

import "fmt"

type Checkout struct {
	TotalPrice     float64
	CurrencySymbol string
}

func (c *Checkout) String() string {
	return fmt.Sprintf("%s%.2f", c.CurrencySymbol, c.TotalPrice)
}
