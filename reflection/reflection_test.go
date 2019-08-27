package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Max", "Dundee"},
			[]string{"Max", "Dundee"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Max", 25},
			[]string{"Max"},
		},
		{
			"Nested fields",
			Person{
				"Max",
				Profile{25, "Dundee"},
			},
			[]string{"Max", "Dundee"},
		},
		{
			"Pointers to things",
			&Person{
				"Max",
				Profile{25, "Dundee"},
			},
			[]string{"Max", "Dundee"},
		},
		{
			"Slices",
			[]Profile{
				{25, "Dundee"},
				{33, "London"},
			},
			[]string{"Dundee", "London"},
		},
		{
			"Arrays",
			[2]Profile{
				{25, "Dundee"},
				{33, "London"},
			},
			[]string{"Dundee", "London"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contains %q but it didnt", haystack, needle)
	}
}
