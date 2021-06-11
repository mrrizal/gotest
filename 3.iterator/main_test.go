package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Run("test repeat", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		assert.Equal(t, got, want)
	})

	t.Run("test repeat zero n", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""

		assert.Equal(t, got, want)
	})

	t.Run("empty char", func(t *testing.T) {
		got := Repeat("", 3)
		want := ""
		assert.Equal(t, got, want)
	})

	t.Run("n less than 0", func(t *testing.T) {
		got := Repeat("a", -5)
		want := ""
		assert.Equal(t, got, want)
	})
}
