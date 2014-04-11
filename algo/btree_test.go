package algo

import "testing"

import "fmt"

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
	_, root := Insert(T(3), nil)

	fmt.Printf("\nroot %v\n", root)

	Insert(T(4), root)
	Insert(T(2), root)

	r, node := Find(T(4), root)
	//fmt.Printf("\nNode= %v\n", node)
	if !r {
		t.Fail()
	}

	r, node = Find(T(5), root)
	if r || node != nil {
		t.Fail()
	}

	r = Delete(T(4), root)
	if !r {
		t.Fail()
	}

	//r = FindLoop(root, 4)
	//fmt.Printf("\ncontains 4\n", r)

}
