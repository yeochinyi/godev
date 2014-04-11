package algo

import (
//"fmt"
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
	parent *Node
	dir    int
	data   Comparable
	links  [2]*Node
}

//func Traverse()

func Insert(data Comparable, root *Node) (ok bool, n *Node) {
	if root == nil {
		return true, &Node{data: data}
	}
	return insertRecur(data, root, none)
}

func insertRecur(data Comparable, parent *Node, dir int) (bool, *Node) {
	if dir != none && parent.links[dir] == nil {
		parent.links[dir] = &Node{parent: parent, dir: dir, data: data}
		return true, parent.links[dir]
	}
	switch parent.data.Compare(data) {
	case equalTo:
		return false, parent
	case greaterThan:
		dir = right
	case lessThan:
		dir = left
	}
	return insertRecur(data, parent, dir)
}

func Find(data Comparable, root *Node) (ok bool, n *Node) {
	return findRecur(data, root)
}

func findRecur(data Comparable, node *Node) (bool, *Node) {
	if node == nil {
		return false, nil
	}
	var dir int
	switch node.data.Compare(data) {
	case equalTo:
		return true, node
	case greaterThan:
		dir = right
	case lessThan:
		dir = left
	}
	return findRecur(data, node.links[dir])
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
	case 2:
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
