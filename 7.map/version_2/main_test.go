package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		assert.Equal(t, got, want)
		assert.Equal(t, err, nil)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := ""
		assert.Equal(t, got, want)
		assert.NotEqual(t, err, nil)
	})

}
