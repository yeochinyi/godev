package main

import "fmt"
import "container/heap"
import "time"

//import "strings"
import "github.com/yeochinyi/godev/json"

type Equaler interface {
	Equal(Equaler) bool
}

type T1 int

func (t T1) Equal(u Equaler) bool {
	return t == u.(T1)
}

type T2 struct {
	value int
}

func (t *T2) Equal(u Equaler) bool {
	return *t == *u.(*T2)
}

func main123() {
	t1 := T1(1)
	b1 := t1.Equal(T1(1))
	fmt.Printf("T1 result is %v\n", b1)

	t2 := &T2{value: 1}
	b2 := t2.Equal(&T2{value: 1})
	fmt.Printf("T2 result is %v\n", b2)
}

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func main123123() {

	clientRequests := make(chan *Request, 3)

	//go handle(clientRequests)
	go func() {
		for {
			select {
			case r := <-clientRequests:
				r.resultChan <- r.f(r.args)
			}
		}
	}()

	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// Send request
	clientRequests <- request
	// Wait for response.
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main123000() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

func main1111() {
	a1 := []int{2, 3, 4}
	//a2 := []int{1}
	a := append([]int{1}, a1...)
	fmt.Printf("%v", a)
}

type IceCreamMaker interface {
	// Hello greets a customer
	Hello()
}

type Ben struct {
	name string
}

func (b *Ben) Hello() {
	//if b.name !=
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
}

type Jerry struct {
	name string
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
}

func main() {
	json.Start()
}

func main1111111() {
	var ben = &Ben{"Ben"}
	var jerry = &Jerry{"Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for {
		time.Sleep(1000)
		maker.Hello()

	}
}
