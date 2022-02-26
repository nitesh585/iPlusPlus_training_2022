#### Introduction

`The Adapter Pattern` is responsible for adaptation of two incompatible
interfaces. It is a structural pattern that is responsible to join
functionalities of independent or incompatible interfaces without modifing
their implementation.

Interfaces may be incompatible but the inner functionality should suit the
need. It allows otherwise incompatible objects to work together by converting
the interface of each struct into an interface expected by the clients.

#### Purpose

- Impedance match an old component to a new system
- Wrap the interface of a object into another interface clients expect.
- Adapter lets objects work together, that could not otherwise because of incompatible interfaces.

#### Design Pattern Diagram

The structs/objects that are participating in adapter pattern are illustrated
in the following diagram:

![alt tag](http://blog.ralch.com/media/golang/design-patterns/adapter.gif)

- `Target` is the domain-specific interface that Client wants to use.
- `Adapter` adapts the interface `Adaptee` to the `Target` interface. It
	implements the `Target` interface in terms of the Adaptee.
- `Adaptee` defines an existing interface that needs adapting.
- `Client` collaborates with objects conforming to the `Target` interface.

The `Target` interface enables objects of adaptee types to be interchangeable
with any other objects that might implement the same interface. However, the
adaptees might not conform to the `Target`. The interface alone is not a
sufficiently powerful mechanism. We need the Adapter pattern. An `Adaptee`
offers similar functionality to the client, but under a different name and with
possibly different parameters. The `Adaptee` is completely independent of the
other classes and is oblivious to any naming conventions or signatures that
they have.

#### Implementation

Lets explore how we should use the Adapter Design Pattern to adopt two
incompatible payment systems and make them available for our customers. Assume
that we are building system that should support `PayPal` and `Bank` payments.
In addition, we are consuming two external libraries that handles each of this
payment methods.

```Golang
package paypal

import (
	"errors"
	"fmt"
	"regexp"
)

var mailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

// Money of PayPal transactions
type Money struct {
	// Amount
	Amount float64
	// Currency for that amount
	Currency string
}

// Payment in PayPal
type Payment struct {
	// APIKey is the PayPal API key
	APIKey string
}

// Send money
func (*Payment) Send(senderEmail, recipientEmail string, money *Money) error {
	if !mailRegexp.MatchString(senderEmail) {
		return errors.New("Invalid sender email address")
	}

	if !mailRegexp.MatchString(recipientEmail) {
		return errors.New("Invalid recipient email address")
	}

	if money == nil {
		return errors.New("The money must be provided")
	}

	if money.Amount <= 0 {
		return errors.New("The amount cannot be negative")
	}

	if money.Currency == "" {
		return errors.New("The currency must be provided")
	}

	fmt.Printf("Send %f %s from %s to %s", money.Amount, money.Currency, senderEmail, recipientEmail)
	return nil
}
```

```Golang
package bank

import (
	"errors"
	"fmt"
	"time"
)

// AccountType determines the type of bank account
type AccountType uint8

const (
	// AccountTypeCurrent is a current bank account
	AccountTypeCurrent AccountType = iota
	// AccountTypeSaving is a saving bank account
	AccountTypeSaving
)

// Account is a bank account
type Account struct {
	// Owner is the bank account owner
	Owner string
	// Email of the owner
	Email string
	// Balance is the bank account balance
	Balance float64
	// Currency of the account
	Currency string
}

// Transaction is the bank transaction
type Transaction struct {
	FromAccount *Account
	ToAccount   *Account
	Amount      float64
	Date        time.Time
	Reason      string
}

// Gateway for the Bank
type Gateway struct {
	// Token Key
	Token string
	// Accounts
	Accounts []*Account
}

// FindAccountByEmail finds a bank account
func (g *Gateway) FindAccountByEmail(email string) (*Account, error) {
	for _, account := range g.Accounts {
		if account.Email == email {
			return account, nil
		}
	}
	return nil, errors.New("Account Not Found")
}

// ProcessTransaction processes a bank transaction
func (g *Gateway) ProcessTransaction(t *Transaction) error {
	if t.FromAccount == nil {
		return errors.New("FromAccount is missing")
	}
	if t.ToAccount == nil {
		return errors.New("ToAccount is missing")
	}

	if t.Reason == "" {
		return errors.New("Reason is not provided")
	}

	if t.Amount <= 0 {
		return errors.New("Invalid amount")
	}

	if t.Amount > t.FromAccount.Balance {
		return errors.New("Insufficient funds")
	}

	fmt.Printf("Transfered %f %s from %s to %s at %v", t.Amount,
		t.FromAccount.Currency, t.FromAccount.Owner, t.ToAccount.Owner, t.Date)

	t.FromAccount.Balance -= t.Amount
	return nil
}
```

We are developing a shopping card that should work with different payment methods:

```Golang
// Checkouter checkouts order
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
```

As you might see, the Bank API and PayPal API cannot be used as different
payment options in the `ShoppingCard` object due to their different signatures.

In order to adopt them we should implement an adapters that obey the `Payment`
interface.

The `BankAdapter` adapts the bank package API by wrapping `bank.Gateway`
struct:

```Golang
// BankAdapter adapts bank API
type BankAdapter struct {
	// Gateway of the bank
	Gateway *bank.Gateway
}

// Pay from email to email this amount
func (b *BankAddapter) Pay(fromEmail, toEmail string, amount float64) error {
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
```

`PayPal` API is adopted by `PayPalAdapter` struct:

```Golang
// PayPalAdapter adapts PayPal API
type PayPalAdapter struct {
	Payment *paypal.Payment
}

// Pay from email to email this amount
func (p *PayPalAdapter) Pay(fromEmail, toEmail string, amount float64) error {
	return p.Payment.Send(fromEmail, toEmail, &paypal.Money{Amount: amount, Currency: "USD"})
}
```

#### Verdict

`The Adapter Pattern` is used wherever there is code to be wrapped up and
redirected to a different implementation.

> But How Much the `Adapter` Should Do?

If the `Target` and `Adaptee` has similarities then the adapter has just to
delegate the requests from the Target to the Adaptee. If `Target` and `Adaptee`
are not similar, then the adapter might have to convert the data structures
between those and to implement the operations required by the Target but not
implemented by the Adaptee.
