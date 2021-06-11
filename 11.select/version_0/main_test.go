package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	assert.Equal(t, want, got)
}
