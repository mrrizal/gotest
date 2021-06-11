package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := Hello("Rizal", "")
		want := "Hello, Rizal"

		assert.Equal(t, got, want)
	})

	t.Run("hello empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assert.Equal(t, got, want)
	})

	t.Run("hello to people lowercase", func(t *testing.T) {
		got := Hello("rizal", "")
		want := "Hello, Rizal"

		assert.Equal(t, got, want)
	})

	t.Run("hello to people uppercase", func(t *testing.T) {
		got := Hello("RIZAL", "")
		want := "Hello, Rizal"

		assert.Equal(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Rizal", "Spanish")
		want := "Hola, Rizal"

		assert.Equal(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Rizal", "French")
		want := "Bonjour, Rizal"

		assert.Equal(t, got, want)
	})

}
