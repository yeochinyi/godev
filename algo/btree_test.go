package algo

import "testing"
import "fmt"

func TestAlgo(t *testing.T) {
	root := Insert(nil, 3)
	Insert(root, 4)
	Insert(root, 2)
	fmt.Printf("\nroot %v\n", root)
	r := FindRecur(root, 4)
	fmt.Printf("\ncontains 4\n", r)
	r = FindLoop(root, 4)
	fmt.Printf("\ncontains 4\n", r)

}
