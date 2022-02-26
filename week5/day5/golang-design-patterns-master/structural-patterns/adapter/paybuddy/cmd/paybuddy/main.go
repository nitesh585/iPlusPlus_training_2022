package main

import (
	"fmt"
	"time"

	"github.com/going/toolkit/log"
	"github.com/svett/golang-design-patterns/structural-patterns/adapter/paybuddy/bank"
	"github.com/svett/golang-design-patterns/structural-patterns/adapter/paybuddy/paypal"
)

// Payment checkouts order
type Payment interface {
	// Pay from email to email this amount
	Pay(fromEmail, toEmail string, amount float64) error
}

// Item in the shopping card
type Item struct {
	// Name of the item
	Name string
	// Price of the item
	Price float64
}

// ShoppingCard in online store
type ShoppingCard struct {
	// Items im the ShoppingCard
	Items []*Item
	// PaymentMethod selected
	PaymentMethod Payment
	// ShopEmailAddress address of the shop
	ShopEmailAddress string
}

// Checkout checkouts a shopping card
func (c *ShoppingCard) Checkout(payeeEmail string) error {
	var total float64

	for _, item := range c.Items {
		total += item.Price
	}

	return c.PaymentMethod.Pay(payeeEmail, c.ShopEmailAddress, total)
}

// BankAdapter adapts bank API
type BankAdapter struct {
	// Gateway of the bank
	Gateway *bank.Gateway
}

// Pay from email to email this amount
func (b *BankAdapter) Pay(fromEmail, toEmail string, amount float64) error {
	fromAccount, err := b.Gateway.FindAccountByEmail(fromEmail)
	if err != nil {
		return err
	}

	toAccount, err := b.Gateway.FindAccountByEmail(toEmail)
	if err != nil {
		return err
	}

	t := &bank.Transaction{
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		Date:        time.Now(),
		Reason:      "Payment to Online Store",
	}

	return b.Gateway.ProcessTransaction(t)
}

// PayPalAdapter adapts PayPal API
type PayPalAdapter struct {
	Payment *paypal.Payment
}

// Pay from email to email this amount
func (p *PayPalAdapter) Pay(fromEmail, toEmail string, amount float64) error {
	return p.Payment.Send(fromEmail, toEmail, &paypal.Money{Amount: amount, Currency: "USD"})
}

func main() {
	payPalAdapter := &PayPalAdapter{
		Payment: &paypal.Payment{
			APIKey: "pay-paly-api-key",
		},
	}

	bankAdapter := &BankAdapter{
		Gateway: &bank.Gateway{
			Token: "bank-token",
			Accounts: []*bank.Account{
				&bank.Account{
					Owner:    "iShop",
					Email:    "shop@example.com",
					Balance:  1000000,
					Currency: "USD",
				},
				&bank.Account{
					Owner:    "Mike Meadows",
					Email:    "mike@example.com",
					Balance:  890300,
					Currency: "USD",
				},
			},
		},
	}

	card := &ShoppingCard{
		Items: []*Item{
			&Item{
				Name:  "Tablet",
				Price: 1000,
			},
			&Item{
				Name:  "Headphones",
				Price: 50,
			},
			&Item{
				Name:  "Smart Watch",
				Price: 550,
			},
		},
		ShopEmailAddress: "shop@example.com",
	}

	fmt.Println("PayPal transaction")
	card.PaymentMethod = payPalAdapter
	if err := card.Checkout("ben.johnson@example.com"); err != nil {
		log.Error(err)
	}

	fmt.Println()

	fmt.Println("Bank transaction")
	card.PaymentMethod = bankAdapter
	if err := card.Checkout("mike@example.com"); err != nil {
		log.Error(err)
	}
}
