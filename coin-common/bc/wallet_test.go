// Package bc
// @Author twilikiss 2024/8/3 1:01:01
package bc

import (
	"fmt"
	"testing"
)

func TestWallet_GetAddress(t *testing.T) {
	wallet, err := NewWallet()
	if err != nil {
		t.Error(err)
	}
	address01 := wallet.GetTestAddress()
	fmt.Println(string(address01))

	priKey := wallet.GetPriKey()
	fmt.Println(priKey)

	err = wallet.ResetPriKey(priKey)
	if err != nil {
		t.Error(err)
	}
	address02 := wallet.GetTestAddress()
	fmt.Println(string(address02))
}
