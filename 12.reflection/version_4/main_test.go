package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		}, {
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		}, {
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		}, {
			"struct with nested field",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		}, {
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		}, {
			"slices",
			[]Profile{Profile{33, "London"}, Profile{33, "Singapore"}},
			[]string{"London", "Singapore"},
		}, {
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assert.Equal(t, true, reflect.DeepEqual(test.ExpectedCalls, got))
		})
	}

	test := struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		"Maps",
		map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		},
		[]string{"Bar", "Boz"},
	}

	t.Run(test.Name, func(t *testing.T) {
		var got []string
		walk(test.Input, func(input string) {
			got = append(got, input)
		})

		for _, want := range test.ExpectedCalls {
			assert.Contains(t, got, want)
		}
	})
}
