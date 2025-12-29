package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	assertExpectedSum(t, sum, expected)
}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertExpectedSum(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}
