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
	first1 = iota
	random2
)

type Node struct {
	data int32
	link [2]*Node
}

func Insert(root *Node, data int32) *Node {
	if root == nil {
		root = &Node{data: data}
	} else if root.data == data {
		return root
	} else {
		dir := 0
		if root.data > data {
			dir = 1
		}
		root.link[dir] = Insert(root.link[dir], data)
	}
	return root
}

func FindRecur(root *Node, data int32) bool {
	if root == nil {
		return false
	} else if root.data == data {
		return true
	} else {
		dir := 0
		if root.data > data {
			dir = 1
		}
		return FindRecur(root.link[dir], data)
	}
}

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
