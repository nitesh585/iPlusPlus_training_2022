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
