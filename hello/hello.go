package main

import (
	"fmt"
	//"time"
	//"godev/newmath"
	//"github.com/yeochinyi/newmath"
	// "net/http/httputil"
	//"bufio"
	//"net"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"io"
	//"encoding/binary"
	//"encoding/hex"
	"strings"
)

func main() {
	s := "/11/222/3333/"
	//strings.
	fmt.Printf("\n%v\n", strings.Join(strings.Split(s, "/")[2:], "/"))
}
