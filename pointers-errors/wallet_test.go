package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit currency", func(t *testing.T) {
		startingBalance := Bitcoin(0)
		wallet := Wallet{startingBalance}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw currency", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(90))
	})

	t.Run("overdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(110))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, expected Bitcoin) {
	t.Helper()

	balance := w.Balance()

	if balance != expected {
		t.Errorf("got %s, but expected %s", balance, expected)
	}
}

func assertError(t testing.TB, err error, expectedError error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected error")
	}

	if err != expectedError {
		t.Errorf("got %s, but expected %s", err.Error(), expectedError.Error())
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("error not expected")
	}
}
