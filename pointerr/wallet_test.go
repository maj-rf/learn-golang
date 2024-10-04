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
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
