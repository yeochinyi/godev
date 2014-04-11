package main

import "fmt"

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

func main() {

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
