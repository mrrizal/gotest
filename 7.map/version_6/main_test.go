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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("new word", func(t *testing.T) {
		err := dictionary.Add("test", "this is just a test")
		assert.Equal(t, err, nil)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Add("test", "this is just a test")
		assert.NotEqual(t, err, nil)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("existing word", func(t *testing.T) {
		err := dictionary.Update("test", "new test")
		assert.Equal(t, err, nil)
	})

	t.Run("not exist word", func(t *testing.T) {
		err := dictionary.Update("foo", "bar")
		assert.NotEqual(t, err, nil)
	})

	t.Run("test search", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "new test"
		assert.Equal(t, err, nil)
		assert.Equal(t, got, want)

		got, err = dictionary.Search("foo")
		want = ""
		assert.NotEqual(t, err, nil)
		assert.Equal(t, got, want)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("delete existing word", func(t *testing.T) {
		err := dictionary.Delete("test")
		assert.Equal(t, err, nil)

		_, err = dictionary.Search("test")
		assert.Equal(t, err, ErrNotFound)
	})

	t.Run("delete unexisting word", func(t *testing.T) {
		err := dictionary.Delete("test")
		assert.Equal(t, err, ErrNotFound)
	})
}