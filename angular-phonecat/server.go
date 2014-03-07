package main

import (
	"fmt"
	//	"html/template"
	"io/ioutil"
	"net/http"
	//"os"
	//"regexp"
	"encoding/json"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"strings"
)

type Phone struct {
	Age      int32
	Id       int32
	ImageUrl string
	Name     string
	Snippet  string
}

func main() {
	d, _ := ioutil.ReadFile("app/phones/phones.json")
	var phones []Phone
	json.Unmarshal(d, &phones)
	fmt.Printf("%v", len(phones))

	session, _ := mgo.Dial("localhost")
	defer session.Close()
	c := session.DB("test").C("phones")
	for _, p := range phones {
		err := c.Insert(p)
		if err != nil {
			panic(err)
		}
	}

}

func start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		listHandler(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {

	file := strings.TrimLeft(r.URL.Path, "/")
	f, err := ioutil.ReadFile("app/" + file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(f)

}
