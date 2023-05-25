package integers //using the package integers to group all of the integer manipulation functions together
//packages should be lowercase because there are still OS' that can't deal with mixed cases lol

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// You can have Examples as part of the test suite
// if you wanted to see which tests are being executed use the -v tag
// Examples won't run without the //Output: 6 comment
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}
