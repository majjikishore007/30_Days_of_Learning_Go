package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 4)
	want := 6
	if got != want {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
