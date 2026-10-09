package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %v", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	// Output: aaaaa
}
