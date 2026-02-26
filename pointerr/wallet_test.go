package pointerr

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		// usamos %s para que se aplique la función String()
		// para Bitcoin
		if got != want {
			t.Errorf("got: %s and want: %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		stringBalance := Bitcoin(20)
		wallet := Wallet{stringBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, stringBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
