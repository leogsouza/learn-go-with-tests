package pointers

// Wallet represents a wallet of bitcoins
type Wallet struct {
	balance int
}

// Deposit increase amount of cryptocurrency
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns total of amount
func (w *Wallet) Balance() int {
	return w.balance
}
