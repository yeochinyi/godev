package algo

import "testing"

import "fmt"

type T int

func (c T) Compare(o Comparable) int {
	oo := o.(T)
	if c == oo {
		return equalTo
	} else if oo > c {
		return greaterThan
	}
	return lessThan
}

func TestAlgo(t *testing.T) {
	_, root := Insert(T(5), nil)

	setRecursive = false

	//fmt.Printf("\nroot %v\n", root)
	for _, x := range []int{1, 4, 2, 6, 15, 9, 17, 12, 20, 10} {
		Insert(T(x), root)
	}

	Traverse(root)

	fmt.Println("\nTest Find 20.")
	r, n := Find(T(20), root)
	if !r || n.data != T(20) {
		t.FailNow()
	}

	fmt.Println("\nTest can't find 100.")
	r, n = Find(T(100), root)
	if r || n != nil {
		t.FailNow()
	}

	fmt.Println("\nGetNumChild(root)")
	c := GetNumChild(root)
	if c != 2 {
		t.FailNow()
	}

	fmt.Println("\nDelete 5")
	r = Delete(T(5), root)
	if !r {
		t.FailNow()
	}

	fmt.Println("\nFind 5")
	r, n = Find(T(5), root)
	if r || n != nil {
		t.FailNow()
	}

	//fmt.Print("\n\n\n")
	//Traverse(root)

	//r = FindLoop(root, 4)
	//fmt.Printf("\ncontains 4\n", r)

}
