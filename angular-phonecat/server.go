package main

import (
	"fmt"
	//	"html/template"
	"io/ioutil"
	"net/http"
	//"os"
	//"regexp"
	"encoding/json"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
	//"strings"
)

const (
	Address string = ":9090"
)

/*
type Phone struct {
	Age      int64
	Id       string
	ImageUrl string
	Name     string
	Snippet  string
}
*/

var session *mgo.Session
var db *mgo.Database

func main() {

	session, _ = mgo.Dial("localhost")
	db = session.DB("test")
	defer session.Close()

	//start()
	importMongo()
}

func importMongo() {
	d, _ := ioutil.ReadFile("app/phones/phones.json")
	//var phones []Phone
	var newPhones []map[string]interface{}
	json.Unmarshal(d, &newPhones)
	fmt.Printf("\n%v", len(newPhones))
	//fmt.Printf("\n%v", newPhones)

	_, mapOfMap := findAllPhones()
	//var oldPhones []map[string]interface{}

	for _, p := range newPhones {
		if str, ok := p["id"].(string); ok {
			if mapOfMap[str] != nil {
				//fmt.Printf("%v already exists", str)
				p = mapOfMap[str]
			} else {
				fmt.Printf("Inserting %v", str)

				//err := db.C("phonesV2").Insert(p)
				//if err != nil {
				//	panic(err)
				//}

			}

			//fmt.Printf("\n Json phone = %v", p)
			d, _ := ioutil.ReadFile("app/phones/" + str + ".json")
			var details map[string]interface{}
			json.Unmarshal(d, &details)
			fmt.Printf("\n Additional = %v", details)

			for k, v := range p {
				details[k] = v
			}

			_, err := db.C("phonesV2").Upsert(p, details)
			if err != nil {
				panic(err)
			}
		}
	}
}

func findAllPhones() ([]map[string]interface{}, map[string]map[string]interface{}) {
	//var result []Phone
	mapOfMap := make(map[string]map[string]interface{})
	var result []map[string]interface{}

	err := db.C("phonesV2").Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	for _, m := range result {
		if str, ok := m["id"].(string); ok {
			fmt.Printf("\nstr = %v", str)
			mapOfMap[str] = m
		}
	}

	//fmt.Printf("\n%v", string(test))
	return result, mapOfMap
}

func start() {
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			listHandler(w, r)
		})
		http.ListenAndServe(":8080", nil)
	*/

	/*
		goweb.Map("/", func(c context.Context) error {

			file := strings.TrimLeft(c.HttpRequest().URL.Path, "/")
			fmt.Printf("\ndebug: file=%v", file)
			f, err := ioutil.ReadFile("app/" + file)
			if err != nil {
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				return goweb.Respond.With(c, http.StatusInternalServerError, []byte("can't get file."))
			}
			return goweb.Respond.With(c, 200, []byte(f))
		})
	*/

	goweb.Map("GET", "phones/phones.json", func(c context.Context) error {
		mapArray, _ := findAllPhones()
		jsonBin, _ := json.Marshal(mapArray)
		return goweb.API.RespondWithData(c, string(jsonBin))

	})

	goweb.MapStatic("/", "app/")

	goweb.Map(func(c context.Context) error {

		// just return a 404 message
		return goweb.API.Respond(c, 404, nil, []string{"File not found"})

	})

	http.ListenAndServe(Address, goweb.DefaultHttpHandler())
}

/*
func listHandler(w http.ResponseWriter, r *http.Request) {

	file := strings.TrimLeft(r.URL.Path, "/")
	f, err := ioutil.ReadFile("app/" + file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(f)
}
*/
