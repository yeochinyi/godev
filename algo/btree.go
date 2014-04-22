package algo

import (
	"fmt"
)

const (
	lessThan = iota
	equalTo
	greaterThan
)

const (
	left = iota
	right
	none
)

type Comparable interface {
	Compare(c Comparable) int
}

type Node struct {
	parent *Node //For easy recur to return parent when found
	dir    int   //left or right
	data   Comparable
	links  [2]*Node
}

var setRecursive bool

func Traverse(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf(" %v ", node.data)
	Traverse(node.links[left])

	Traverse(node.links[right])

	//fmt.Println("")
}

func traverseRecur(node *Node, dir int) {
	if node == nil {
		return
	}
	fmt.Printf(" %v ", node.data)
	Traverse(node.links[left])
	Traverse(node.links[right])

	fmt.Println("")
}

func Insert(data Comparable, root *Node) (ok bool, newNode *Node) {
	if root == nil {
		return true, &Node{data: data}
	}
	if setRecursive {
		return insertRecur(data, root, none)
	} else {
		return insertLoop(data, root)
	}

}

func insertLoop(data Comparable, parent *Node) (ok bool, newNode *Node) {

	curr := parent
	var dir int

	for curr != nil {
		//fmt.Printf("\n data is %v", curr.data)
		switch curr.data.Compare(data) {
		case equalTo:
			return false, nil
		case greaterThan:
			dir = right
		case lessThan:
			dir = left
		}
		parent = curr
		curr = curr.links[dir]
	}

	parent.links[dir] = &Node{parent: parent, dir: dir, data: data}

	return true, parent.links[dir]
}

//Recurse until empty spot or return equalTo found
func insertRecur(data Comparable, parent *Node, dir int) (ok bool, newNode *Node) {
	switch parent.data.Compare(data) {
	case equalTo:
		return false, parent
	case greaterThan:
		dir = right
	case lessThan:
		dir = left
	}

	if parent.links[dir] == nil { //Found insertion point
		parent.links[dir] = &Node{parent: parent, dir: dir, data: data}
		return true, parent.links[dir]
	}

	return insertRecur(data, parent.links[dir], dir)
}

func Find(data Comparable, root *Node) (ok bool, theOne *Node) {
	if setRecursive {
		return findRecur(data, root)
	} else {
		//fmt.Println("Using findLoop.")
		return findLoop(data, root)
	}
}

func findLoop(data Comparable, parent *Node) (ok bool, theOne *Node) {
	curr := parent
	for curr != nil {
		//fmt.Printf("\n data is %v", curr.data)
		var dir int
		switch curr.data.Compare(data) {
		case equalTo:
			return true, curr
		case greaterThan:
			dir = right
		case lessThan:
			dir = left
		}
		curr = curr.links[dir]
	}
	return false, nil
}

func findRecur(data Comparable, parent *Node) (ok bool, theOne *Node) {
	if parent == nil {
		return false, nil
	}
	var dir int
	switch parent.data.Compare(data) {
	case equalTo:
		return true, parent
	case greaterThan:
		dir = right
	case lessThan:
		dir = left
	}
	return findRecur(data, parent.links[dir])
}

func GetNumChild(node *Node) int {
	count := 0
	for _, x := range node.links {
		if x != nil {
			count++
		}
	}
	return count
}

func Delete(data Comparable, root *Node) (ok bool) {
	found, node := Find(data, root)
	if !found {
		return false
	}
	switch GetNumChild(node) {
	case 0:
		node.parent.links[node.dir] = nil
		return true
	case 1:
		nonEmpty := left
		if node.links[left] == nil {
			nonEmpty = right
		}
		node.parent.links[node.dir] = node.links[nonEmpty]
		node.links[nonEmpty] = nil
	case 2: //traverse 1 left and right all e way
		curr := node.links[left]
		for ; curr.links[right] != nil; curr = curr.links[right] {
		}
		node.data = curr.data //swap data
		//if last node has left link.. reconnect to parent
		if curr.links[left] != nil {
			curr.parent.links[right] = curr.links[left]
		} else {
			//nullifed parent link
			curr.parent.links[left] = nil
		}
	}

	return true
}

/*
func FindLoop(root *Node, data int32) bool {

	for root != nil {
		if root.data == data {
			return true
		} else {
			dir := 0
			if root.data > data {
				dir = 1
			}
			root = root.link[dir]
		}
	}
	return false
}
*/
