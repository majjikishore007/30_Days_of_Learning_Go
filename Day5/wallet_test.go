package day5

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}
	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		got := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, got, ErrInsufficientFunds.Error())
	})
}
