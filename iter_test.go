package iter

import "testing"

func TestIter(t *testing.T) {
	for i := range R(20) {
		println(i)
	}
}
