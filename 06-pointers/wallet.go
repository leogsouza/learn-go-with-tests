package pointers

import (
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

// Deposit increase amount of cryptocurrency
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns total of amount
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
