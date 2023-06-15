package main

import (
	"errors"

	"github.com/Serhii1Epam/oopDesignPatterns/internal/pkg/strategy"
)

func ProcPayment(p strategy.Payment) error {
	err := p.Pay()
	if err != nil {
		return errors.New("Payment error!")
	}
	return err
}

func main() {
	payType := 3

	var pay strategy.Payment

	switch payType {
	case 1:
		{
			pay = strategy.NewCardPayment("5453...4433", 234)
		}
	case 2:
		{
			pay = strategy.NewCryptoPayment([]byte("4f5d9c1a"), strategy.USDT)
		}
	case 3:
		{
			pay = strategy.NewPayPalPayment("https://paypal.com/...")
		}
	default:
		{
			return
		}
	}
	ProcPayment(pay)
}
