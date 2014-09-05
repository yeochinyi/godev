package json

import (
	"fmt"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"strings"
	"testing"
	"time"
)

/*
func Test(t *testing.T) {
	session, _ := mgo.Dial("test:test@localhost")
	db = session.DB("test")
	search := bson.M{"_id": bson.ObjectIdHex("52f499bc616ee7e25bb662ee")}
	//search := bson.M{"name": "Ale"}

	var result []map[string]interface{}
	err := db.C("people").Find(search).All(&result)
	fmt.Printf("\nResults:= %v", result)
	if err != nil {
		fmt.Printf("\nError in get people due to %v", err)
	}
}
*/

type BArray struct {
	b []byte
}

var client *http.Client

func TestServer(t *testing.T) {

	client = &http.Client{}

	fmt.Println("\nStarting....")
	go Start()
	time.Sleep(1 * time.Second)
	fmt.Println("\nStarted.")

	var b []byte
	Do("c", t, &b)
	Do("c/people", t, &b)

	var m []map[string]interface{}
	json.Unmarshal(b, &m)
	name := m[0]["name"].(string)
	Do("c/people?name="+name, t, &b)
	id := m[0]["_id"].(string)
	Do("c/people/"+id, t, &b)
	//Post("c/people/"+id+"?name=changed", t, &b)

}

func Do(part string, t *testing.T, b *[]byte) {

	//r, err := http.Get("http://127.0.0.1:9090/rest/" + part)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:9090/rest/"+part, nil)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	*b, _ = ioutil.ReadAll(resp.Body)

	s := string(*b)
	fmt.Printf("\n%v:->%v", part, s)

	if s == "" {
		t.Fail()
	}

}
