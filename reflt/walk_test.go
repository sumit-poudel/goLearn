package reflt

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
			"struct with one string field",
			struct {
				Name string
			}{"Sumit"},
			[]string{"Sumit"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Sumit", "Chitwan"},
			[]string{"Sumit", "Chitwan"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Sumit", 20},
			[]string{"Sumit"},
		},
		{
			"nested struct",
			Person{"Sumit", Profile{20, "chitwan"}},
			[]string{"Sumit", "chitwan"},
		},
		{
			"pointers to things",
			&Person{"Sumit", Profile{20, "chitwan"}},
			[]string{"Sumit", "chitwan"},
		},
		{
			"slices",
			[]Profile{{20, "chitwan"}, {30, "test"}},
			[]string{"Chitwan", "test"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}
}
