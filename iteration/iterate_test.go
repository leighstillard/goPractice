package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	const c int = 10

	repeated := Repeat("a", c)
	expected := ""
	for i := 0; i < c; i++ {
		expected += "a"
	}

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	const c int = 10

	for i := 0; i < b.N; i++ {
		Repeat("a", c)
	}
}
func ExampleRepeat() {
	output := Repeat("z", 10)
	fmt.Println(output)
	// Output: zzzzzzzzzz
}
