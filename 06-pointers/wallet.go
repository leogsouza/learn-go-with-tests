package pointers

import (
	"errors"
	"fmt"
)

// Biticoin represents the amount of bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a wallet of bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit increase the amount of Bitcoins
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns total of amount
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw decrease the amount of Bitcoins
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
