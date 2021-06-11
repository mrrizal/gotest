package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGree(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Rizal")

	got := buffer.String()
	want := "Hello, Rizal"

	assert.Equal(t, want, got)
}
