// Behavioural pattern Strategy
package strategy

import (
	"fmt"
)

type Payment interface {
	Pay() error
}

const (
	ETH  = "Ethirium"
	BTC  = "Bitcoin"
	USDT = "TRC20 USD"
)

/*func (e *expiryDate) YearAndMonth() string {
	var output string
	fmt.Printf(output, "%2d/%2d", e.year, e.month)
	return output
}

func (e *expiryDate) MonthAndYear() string {
	var output string
	fmt.Printf(output, "%2d/%2d", e.month, e.year)
	return output
}*/

type cardPayment struct {
	cardNumber string
	cvv        int
}

func NewCardPayment(card string, code int) Payment {
	return &cardPayment{cardNumber: card, cvv: code}
}

func (c *cardPayment) Pay() error {
	// Pay with Card
	fmt.Printf("Payment using cardPayment structure. Card: %s", c.cardNumber)
	return nil
}

type payPalPayment struct {
	payPalApi string
}

func NewPayPalPayment(a string) Payment {
	return &payPalPayment{payPalApi: a}
}

func (pp *payPalPayment) Pay() error {
	//Pay with PayPal
	fmt.Printf("Payment using payPalPayment structure. API: %s", pp.payPalApi)
	return nil
}

type cryptoPayment struct {
	wallet []byte
	net    string
}

func NewCryptoPayment(w []byte, n string) Payment {
	return &cryptoPayment{wallet: w, net: n}
}

func (c cryptoPayment) PrintInHex() {
	fmt.Printf("0x%x", c.wallet)
}

func (c *cryptoPayment) Pay() error {
	//Pay with Crypto
	fmt.Printf("Payment using cryptoPayment structure. Network :%s\n", c.net)
	return nil
}
