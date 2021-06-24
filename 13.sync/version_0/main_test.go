package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Counter{}
	counter.Inc()
	counter.Inc()
	counter.Inc()

	want := 3
	got := counter.Value()

	assert.Equal(t, want, got)
}
