package iteration

import (
	"fmt"
	"testing"
)

func Test_Repeater(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// go test -bench="." to run benchmarks
// Benchmarks essentially allow us to gauge on average how fast our program runs!
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}

func ExampleRepeat() {
	repeat := Repeat("a", 7)
	fmt.Println(repeat)
	//Output: aaaaaaa
}
