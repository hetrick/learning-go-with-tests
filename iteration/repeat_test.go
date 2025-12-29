package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 4 times", func(t *testing.T) {
		assertRepeated(t, Repeat("a", 4), "aaaa")
	})
	t.Run("repeat 2 times", func(t *testing.T) {
		assertRepeated(t, Repeat("b", 2), "bb")
	})
	t.Run("repeat 1 time", func(t *testing.T) {
		assertRepeated(t, Repeat("c", 1), "c")
	})
}

func assertRepeated(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q, but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("z", 7)
	fmt.Println(repeated)
	// Output: zzzzzzz
}
