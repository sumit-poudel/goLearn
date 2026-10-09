package pointerr

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := ProcessWithDrawal(&wallet, "acc-123", 20)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := ProcessWithDrawal(&wallet, "chut", 40)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Wanted: %d, Got %d", want, got)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didnt got an error but wanted on")
	}
	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got error when shouldnt get one")
	}
}
