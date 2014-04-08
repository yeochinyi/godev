package algo

import (
//"fmt"
//"math/rand"
//"time"
//"sync"
//"godev/newmath"
//"github.com/yeochinyi/newmath"
// "net/http/httputil"
//"bufio"
//"net"
//"io/ioutil"
//"net/http"
//"net/url"
//"os"
)

const (
	lessThan = iota
	equalTo
	greaterThan
)

type Comparable interface {
	Compare(c Comparable) int
}

type Node struct {
	data Comparable
	link [2]*Node
}

func InsertRecur(root *Node, data Comparable) *Node {
	if root == nil {
		root = &Node{data: data}
	} else if root.data.Compare(data) == equalTo {
		return root
	} else {
		dir := 0
		if root.data.Compare(data) == greaterThan {
			dir = 1
		}
		root.link[dir] = InsertRecur(root.link[dir], data)
	}
	return root
}

/*
func FindRecur(root *Node, data Comparable) bool {
	if root == nil {
		return false
	} else if root.data.Compare(data) == equalTo {
		return true
	} else {
		dir := 0
		if root.data.Compare(data) == greaterThan {
			dir = 1
		}
		return FindRecur(root.link[dir], data)
	}
}*/

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
