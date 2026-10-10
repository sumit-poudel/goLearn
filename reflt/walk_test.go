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
			[]string{"chitwan", "test"},
		},
		{
			"arrays",
			[2]Profile{{20, "chitwan"}, {30, "test"}},
			[]string{"chitwan", "test"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
		})
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
		assertLength(t, got, len(aMap))
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func assertLength(t testing.TB, got []string, want int) {
	t.Helper()
	if len(got) != want {
		t.Errorf("expected %v, but got %v", got, want)
	}
}
