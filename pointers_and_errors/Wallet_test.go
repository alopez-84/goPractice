package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	WalletTest := func(t testing.TB, wallet Wallet, want IguanaDolla) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("Wanted an Error but did not get one")
		}
	}

	t.Run("testing Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(IguanaDolla(10))

		WalletTest(t, wallet, IguanaDolla(10))
	})

	t.Run("Testing withdraw", func(t *testing.T) {
		wallet := Wallet{balance: IguanaDolla(20)}
		err := wallet.Withdraw(IguanaDolla(10))

		assertNoError(t, err)
		WalletTest(t, wallet, IguanaDolla(10))
	})

	t.Run("Withdraw Insufficient Funds", func(t *testing.T) {
		wallet := Wallet{IguanaDolla(20)}
		err := wallet.Withdraw(IguanaDolla(100))

		assertError(t, err, "cannont withdraw, insufficient funds")
		WalletTest(t, wallet, IguanaDolla(20))

	})
}
