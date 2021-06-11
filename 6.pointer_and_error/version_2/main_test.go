package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assert.Equal(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		assert.Equal(t, got, want)
		assert.Equal(t, err, nil)

	})

	t.Run("Withdrawy less than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(20))

		got := wallet.Balance()
		want := Bitcoin(10)

		assert.Equal(t, got, want)
		assert.NotEqual(t, err, nil)

	})

	t.Run("Withdrawy zero balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0)}

		err := wallet.Withdraw(Bitcoin(20))

		got := wallet.Balance()
		want := Bitcoin(0)

		assert.Equal(t, got, want)
		assert.NotEqual(t, err, nil)

	})
}
