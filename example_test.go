package iter_test

import (
	"fmt"

	"github.com/aerth/iter"
)

func ExampleR() {
	for i := range iter.R(3) {
		fmt.Println(i)
	}
	// Output:
	// 0
	// 1
	// 2
}
