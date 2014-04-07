package algo

import (
	//"fmt"
	"math/rand"
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
	first  = iota
	random = iota
	last   = iota
)

/*
func main() {
	start := time.Now()
	a := []int{5, 4, 3, 2, 1, 2, 3, 4, 5}
	for i := 0; i < 10; i++ {
		a = append(a, a...)
	}
	fmt.Println(a)
	a = sort(a, random)
	fmt.Println(a)
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
	//time.Sleep(3 * time.Second)

}*/

func QuickSort(a []int, idx int) []int {
	//return recurEasy(a, idx)
	return qRecurSmall(a, idx, nil)
}

func qGetIndex(a []int, idx int) int {
	switch idx {
	case first:
		return 1
	case random:
		return rand.Intn(len(a))
	case last:
		return len(a) - 1
	}
	return 0
}

func qRecurEasy(a []int, idx int) []int {
	if len(a) == 0 {
		return a
	}
	index := qGetIndex(a, idx)

	var less, more []int
	for i, v := range a {
		if i == index {
			continue
		} else if v < a[index] {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}
	less = qRecurEasy(less, idx)
	more = qRecurEasy(more, idx)

	return append(append(less, a[index]), more...)
}

func qRecurSmall(a []int, idx int, ch chan int) []int {
	defer func() {
		if ch != nil {
			ch <- 1
		}
		//close(ch)
	}()
	//fmt.Println(a)
	if len(a) < 2 {
		return a
	}
	index := qGetIndex(a, idx)
	left, right := 0, len(a)-1
	a[index], a[right] = a[right], a[index]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	//var wg sync.WaitGroup
	ch1 := make(chan int)
	go qRecurSmall(a[:left], idx, ch1)
	//go recurSmall(a[left+1:], idx, ch1)
	//recurSmall(a[:left], idx, nil)
	qRecurSmall(a[left+1:], idx, nil)
	//<-ch1
	<-ch1
	close(ch1)
	return a
}
