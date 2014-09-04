package newmath

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {

	fmt.Println("123")
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
