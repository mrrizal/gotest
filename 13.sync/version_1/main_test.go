package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("it run safely concurrently", func(t *testing.T) {
		want := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assert.Equal(t, want, counter.Value())
	})
}
