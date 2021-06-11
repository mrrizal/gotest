package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestSum(t *testing.T) {
	t.Run("test sum", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 6
		assert.Equal(t, got, want)
	})

	t.Run("minus value", func(t *testing.T) {
		got := Sum([]int{-3, 3})
		want := 0
		assert.Equal(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test sum all", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assert.DeepEqual(t, got, want)
		assert.Equal(t, len(got), len(want))
	})

	t.Run("test sum all", func(t *testing.T) {
		got := SumAll([]int{1, 2, 2}, []int{0, 9})
		want := []int{5, 9}

		assert.DeepEqual(t, got, want)
		assert.Equal(t, len(got), len(want))
	})

	t.Run("test sum all zero value", func(t *testing.T) {
		got := SumAll([]int{}, []int{0, 9})
		want := []int{0, 9}

		assert.DeepEqual(t, got, want)
		assert.Equal(t, len(got), len(want))
	})
}

func TestSumAllTrails(t *testing.T) {
	t.Run("test sum all trails", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assert.DeepEqual(t, got, want)
	})

	t.Run("test sum all trails", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2, 5}, []int{0, 9})
		want := []int{7, 9}
		assert.DeepEqual(t, got, want)
	})

	t.Run("test sum all trails", func(t *testing.T) {
		got := SumAllTrails([]int{1}, []int{0, 9})
		want := []int{0, 9}
		assert.DeepEqual(t, got, want)
	})

	t.Run("test sum all trails", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{0, 9})
		want := []int{0, 9}
		assert.DeepEqual(t, got, want)
	})

	t.Run("test sum all trails, all zeros", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{})
		want := []int{0, 0}
		assert.DeepEqual(t, got, want)
	})
}
