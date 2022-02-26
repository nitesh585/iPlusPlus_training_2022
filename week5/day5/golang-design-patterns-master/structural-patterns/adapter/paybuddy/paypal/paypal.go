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
