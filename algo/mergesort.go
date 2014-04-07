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

/*
const (
	first  = iota
	random = iota
	last   = iota
)*/

/*
func main() {
	start := time.Now()
	a := []int{5, 4, 3, 2, 1, 2, 3, 4, 5}
	//for i := 0; i < 10; i++ {
	//	a = append(a, a...)
	//}
	fmt.Println(a)
	a = sort(a, random)
	fmt.Println(a)
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
	//time.Sleep(3 * time.Second)

}*/

func MergeSort(a []int, idx int) []int {
	return mergerRecurEasy(a)
	//return recurSmall(a, idx, nil)
}

func mergerRecurEasy(a []int) []int {
	//fmt.Printf("%d:%d\n", a, len(a))
	if len(a) < 2 {
		//fmt.Printf("Return %d\n", a)
		return a
	}

	half := len(a) / 2
	t1 := mergerRecurEasy(a[:half])
	t2 := mergerRecurEasy(a[half:])

	var ret []int

	for x, y := 0, 0; ; {

		if x == len(t1) {
			ret = append(ret, t2[y:]...)
			break
		} else if y == len(t2) {
			ret = append(ret, t1[x:]...)
			break
		}

		if t1[x] < t2[y] {
			ret = append(ret, t1[x])
			x++
		} else {
			ret = append(ret, t2[y])
			y++
		}

	}

	return ret
}
