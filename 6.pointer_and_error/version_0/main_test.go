package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	assert.Equal(t, got, want)
}
