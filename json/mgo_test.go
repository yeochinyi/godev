package json

import (
	//	"bytes"
	//	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestMgo(t *testing.T) {

	fmt.Println("\nStarting....")
	Start()
	time.Sleep(2 * time.Second)
	fmt.Println("\nStarted.")
	r, err := http.Get("http://127.0.0.1:9090/rest/cs/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nReading....")
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
	fmt.Println("\nRead finished.")

}
