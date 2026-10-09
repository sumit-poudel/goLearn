package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		language string
		want     string
	}{
		{
			name:     "test with name",
			userName: "sumit",
			language: "",
			want:     "Hello, sumit",
		}, {
			name:     "test with blank",
			userName: "",
			language: "",
			want:     "Hello, World",
		}, {
			name:     "in Spanish",
			userName: "sumit",
			language: "Spanish",
			want:     "Hola, sumit",
		},{
			name:     "in French",
			userName: "sumit",
			language: "French",
			want:     "Bonjour, sumit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.userName, tt.language)
			assertCorrectMessage(t, got, tt.want)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("Wanted %s, Got %s", want, got)
	}
}
