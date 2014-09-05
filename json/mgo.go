package json

import (
	"fmt"
	//	"html/template"
	//"io/ioutil"
	"net/http"
	//"os"
	//"regexp"
	//"code.google.com/p/gorest"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

const (
	Address string = "localhost:9090"
)

func main() {
	Start()
}

var session *mgo.Session
var db *mgo.Database

func Start() {

	session, err := mgo.Dial("test:test@localhost")

	if err != nil {
		panic(err)
	}

	db = session.DB("test")
	defer session.Close()

	/*
		http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
			fs := http.FileServer(http.Dir("app"))
			r.URL.Path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
			fmt.Printf("\nr.URL.Path=%v\n", r.URL.Path)
			fs.ServeHTTP(w, r)
		})
	*/

	fmt.Println("\nStarting Handle")
	http.HandleFunc("/rest/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nr.URL.Path=%v\n", r.URL.Path)
		s := strings.Split(r.URL.Path, "/")
		for i, v := range s {
			fmt.Printf("\n%v = %v\n", i, v)
		}
		r.ParseForm()
		fmt.Printf("\nForm =%v\n", r.Form)

		switch s[2] {
		case "helloworld":
			fmt.Fprintf(w, "Hello World!")
		case "c":
			switch len(s) {
			case 3:
				l, _ := db.CollectionNames()
				m, _ := json.Marshal(l)
				w.Write(m)

			case 4:
				search := make(bson.M)
				for k, v := range r.Form {
					search[k] = v[0]
				}
				var result []map[string]interface{}
				db.C(s[3]).Find(search).All(&result)
				m, _ := json.Marshal(result)
				w.Write(m)

			case 5: //get id
				fields := make(bson.M)
				for k, v := range r.Form {
					fields[k] = v[0]
				}
				id := bson.M{"_id": bson.ObjectIdHex(s[4])}
				var result map[string]interface{}
				switch r.Method {
				case "PUT", "POST": //upsert
					changed, _ := db.C(s[3]).Upsert(id, fields)
					//if new, we need to get back id
					idString := changed.UpsertedId.(string)
					id = bson.M{"_id": bson.ObjectIdHex(idString)}
					db.C(s[3]).Find(id).One(&result)
					m, _ := json.Marshal(result)
					w.Write(m)
				case "DELETE":
					db.C(s[3]).RemoveId(id)
					m, _ := json.Marshal("OK")
					w.Write(m)
				default:
					db.C(s[3]).Find(id).One(&result)
					m, _ := json.Marshal(result)
					w.Write(m)
				}
			}
		}

	})
	http.ListenAndServe(Address, nil)
}
