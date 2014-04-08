package algo

import "testing"

//import "fmt"

type T int

func (c T) Compare(o Comparable) int {
	oo := o.(T)
	if c == oo {
		return equalTo
	} else if c > oo {
		return greaterThan
	}
	return lessThan
}

func TestAlgo(t *testing.T) {
	root := InsertRecur(nil, T(3))
	InsertRecur(root, T(4))
	InsertRecur(root, T(2))
	/*
		fmt.Printf("\nroot %v\n", root)
		r := FindRecur(root, 4)
		fmt.Printf("\ncontains 4\n", r)
		r = FindLoop(root, 4)
		fmt.Printf("\ncontains 4\n", r)
	*/
}
