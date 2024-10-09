package pointerr

import (
	"fmt"
	"testing"
)

/*
%p = base 16 notation like 0x1400000e180
& = get the pointer for a memory address
*/

func TestWaller(t *testing.T) {
	t.Run("check balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		assertBalance(t, wallet, Bitcoin(100))
	})
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(10))
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(90))
	})
	t.Run("Withdraw insufficient amount", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(110))
		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
