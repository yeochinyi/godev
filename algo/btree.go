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

const (
	preOrder = iota
	inOrder
	postOrder
)

type Comparable interface {
	Compare(c Comparable) int
}

type Node struct {
	data  Comparable
	links [2]*Node
}

type DirEntry struct {
	*Node
	dir int
}

type LevelDirEntry struct {
	*Node
	dir   int
	level int
}

var setRecursive bool

func Traverse(node *Node, order int, leftFirst bool) []*LevelDirEntry {
	if setRecursive {
		return traverseRecur(node, 0, none, order, leftFirst)
	} else {
		return traverseLoop(node, order, leftFirst)
	}

}

//To put '1 more' detail to differentiate with LevelDirEntry in the slice
type Processing struct {
	*LevelDirEntry
}

func traverseLoop(node *Node, order int, leftFirst bool) []*LevelDirEntry {

	firstGo := left
	secondGo := right

	if !leftFirst {
		firstGo = right
		secondGo = left
	}

	nodes := []*LevelDirEntry{}
	stack := []interface{}{}

	addStack := func(parent *Node, dir int, level int) {
		if parent.links[dir] != nil {
			p := Processing{&LevelDirEntry{parent.links[dir], dir, level + 1}}
			stack = append(stack, &p)
		}
	}

	dir := none
	level := 0

	//using slice as stack.. due to append()",
	// we need to work from the back instead of the front i.e [0]
	// add last to process element in the front[0], pop the back or add new to the back
	for {
		switch order {
		case preOrder:
			addStack(node, secondGo, level)
			addStack(node, firstGo, level)
			stack = append(stack, &LevelDirEntry{node, dir, level})
		case inOrder:
			addStack(node, secondGo, level)
			stack = append(stack, &LevelDirEntry{node, dir, level})
			addStack(node, firstGo, level)
		default: //postOrder
			stack = append(stack, &LevelDirEntry{node, dir, level})
			addStack(node, secondGo, level)
			addStack(node, firstGo, level)
		}

		//2nd loop to clear Nodes into slice
	ClearNodes:
		for {
			if len(stack) == 0 {
				return nodes
			}
			//pop the back
			end := len(stack) - 1
			i := stack[end]
			stack = stack[:end]

			switch i.(type) {
			case *Processing:
				e := i.(*Processing).LevelDirEntry
				level = e.level
				node = e.Node
				dir = e.dir
				break ClearNodes
			case *LevelDirEntry:
				nodes = append(nodes, i.(*LevelDirEntry))
			}
		}
	}
}

func traverseRecur(node *Node, level int, dir int, order int, leftFirst bool) []*LevelDirEntry {

	firstGo := left
	secondGo := right

	if !leftFirst {
		firstGo = right
		secondGo = left
	}

	if node == nil {
		return nil
	}
	l := LevelDirEntry{node, dir, level}
	e := []*LevelDirEntry{&l}
	e1 := traverseRecur(node.links[firstGo], level+1, firstGo, order, leftFirst)
	e2 := traverseRecur(node.links[secondGo], level+1, secondGo, order, leftFirst)

	switch order {
	case preOrder:
		return append(append(e, e1...), e2...)
	case inOrder:
		return append(append(e1, e...), e2...)
	default: //postOrder
		return append(append(e1, e2...), e...)
	}
}

func PrintTree(node *Node) {

	for _, node := range Traverse(node, inOrder, false) {
		c := ""
		switch node.dir {
		case right:
			c = "/"
		case left:
			c = "\\"
		}
		/*
			if node == nil {
				for i := 0; i < level; i++ {
					fmt.Print("\t")
				}
				fmt.Printf("%v~\n\n", c)
				return
			}*/
		for i := 0; i < node.level; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("%v%v\n\n", c, node.data)
	}

}

func Insert(data Comparable, root *Node, atRoot bool) (ok bool, newNode *Node) {
	if root == nil {
		return true, &Node{data: data}
	}
	if setRecursive {
		if atRoot {
			return insertRootRecur(data, root, none)
		}
		return insertRecur(data, root, none)
	} else {
		return insertLoop(data, root)
	}

}

func insertLoop(data Comparable, parent *Node) (ok bool, newNode *Node) {
	curr := parent
	var dir int

	for curr != nil {
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
	parent.links[dir] = &Node{data: data}

	return true, parent.links[dir]
}

func flip(dir int) int {
	switch dir {
	case left:
		return right
	case right:
		return left
	default:
		return none
	}
}

//Recurse until empty spot or return equalTo found
func insertRootRecur(data Comparable, parent *Node, dir int) (ok bool, newNode *Node) {
	switch parent.data.Compare(data) {
	case equalTo:
		return false, parent
	case greaterThan:
		dir = right
	case lessThan:
		dir = left
	}
	var save *Node

	if parent.links[dir] == nil { //Found insertion point
		save = &Node{data: data}
		ok = true
	} else {
		ok, save = insertRootRecur(data, parent.links[dir], dir)
	}
	//This is simply moving the parent to be the node child
	// and relink the "replaced" child to the parent (where the node is supposed to be link)
	parent.links[dir] = save.links[flip(dir)] //relink node's child under parent
	save.links[flip(dir)] = parent            // link parent as node's child

	return ok, save
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
	var save *Node
	if parent.links[dir] == nil { //Found insertion point
		save = &Node{data: data}
		parent.links[dir] = save //can't do on the return as it will be linked to root eventually
		ok = true
	} else {
		ok, save = insertRecur(data, parent.links[dir], dir)
	}

	return ok, save
}

func Find(data Comparable, root *Node) (ok bool, reversePath []*DirEntry) {
	//if setRecursive {
	//	return findRecur(data, root)
	//} else {
	return findLoop(data, root)
	//}
}

func Reverse(orig []interface{}) (ret []interface{}) {
	l := len(orig)
	ret = make([]interface{}, l)
	for i, x := range orig {
		ret[l-i] = x
	}
	return
}

func reverse(nodes []*DirEntry) {
	for x, y := 0, len(nodes)-1; x < y; x, y = x+1, y-1 {
		nodes[x], nodes[y] = nodes[y], nodes[x]
	}
}

func findLoop(data Comparable, parent *Node) (ok bool, reversePath []*DirEntry) {
	nodes := []*DirEntry{}
	curr := parent
	dir := none
	for curr != nil {
		nodes = append(nodes, &DirEntry{curr, dir})
		switch curr.data.Compare(data) {
		case equalTo:
			dir = none
		case greaterThan:
			dir = right
		case lessThan:
			dir = left
		}
		if dir == none {
			reverse(nodes)
			return true, nodes
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

	found, nodes := Find(data, root)
	if !found {
		return false
	}
	nodePath := nodes[0]
	var parentPath *DirEntry
	if len(nodes) > 1 { // in case it's root
		parentPath = nodes[1]
	}
	switch GetNumChild(nodePath.Node) {
	case 0:
		if parentPath == nil {
			return false // if tree have only 1 node as root,do nothing
		}
		parentPath.Node.links[nodePath.dir] = nil
		return true
	case 1:
		nonEmpty := left
		if nodePath.Node.links[left] == nil {
			nonEmpty = right
		}
		nodePath.Node.data = nodePath.Node.links[nonEmpty].data
		if parentPath != nil {
			parentPath.Node.links[nodePath.dir] = nodePath.Node.links[nonEmpty] //relink
		}
	case 2: //traverse 1 left and right all e way
		curr := nodePath.Node.links[left]
		var parentcurr *Node
		for ; curr.links[right] != nil; curr = curr.links[right] {
			parentcurr = curr
		}
		nodePath.Node.data = curr.data //move data
		if parentcurr == nil {
			nodePath.Node.links[left] = nil // if curr is just left of theNode
		} else {
			parentcurr.links[right] = curr.links[left] //replace curr position
		}

	}

	return true
}
