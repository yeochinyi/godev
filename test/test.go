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

func main() {
	t1 := T1(1)
	b1 := t1.Equal(T1(1))
	fmt.Printf("T1 result is %v\n", b1)

	t2 := &T2{value: 1}
	b2 := t2.Equal(&T2{value: 1})
	fmt.Printf("T2 result is %v\n", b2)
}
