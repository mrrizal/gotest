package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just test"}

	got := Search(dictionary, "test")
	want := "this is just test"

	assert.Equal(t, got, want)
}