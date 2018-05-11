package pointers

import "testing"

func TestWallet(t *testing.T) {

	checkBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		checkBalance(t, wallet, want)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		checkBalance(t, wallet, want)

	})

}
