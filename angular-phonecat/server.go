package main

import (
	"fmt"
	//	"html/template"
	"io/ioutil"
	"net/http"
	//"os"
	//"regexp"
	"encoding/json"
	//"github.com/stretchr/goweb"
	//"github.com/stretchr/goweb/context"
	"code.google.com/p/gorest"
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
	"strings"
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

type PhoneService struct {
	//Service level config
	gorest.RestService `root:"/rest/phone/" consumes:"application/json" produces:"application/json"`

	//End-Point level configs: Field names must be the same as the corresponding method names,
	// but not-exported (starts with lowercase)

	listPhones gorest.EndPoint `method:"GET" path:"/phones/" output:"[]map[string]interface {}"`
	getPhone   gorest.EndPoint `method:"GET" path:"/{Id:string}" output:"map[string]interface {}"`

	//addItem     gorest.EndPoint `method:"POST" path:"/items/" postdata:"Item"`

	//On a real app for placeOrder below, the POST URL would probably be just /orders/, this is just to
	// demo the ability of mixing post-data parameters with URL mapped parameters.
	//placeOrder  gorest.EndPoint `method:"POST" path:"/orders/new/{UserId:int}/{RequestDiscount:bool}" postdata:"Order"`
	//viewOrder     gorest.EndPoint `method:"GET" path:"/orders/{OrderId:int}" output:"Order"`
	//deleteOrder     gorest.EndPoint `method:"DELETE" path:"/orders/{OrderId:int}"`

}

func (serv PhoneService) ListPhones() []map[string]interface{} {
	//serv.ResponseBuilder().CacheMaxAge(60*60*24) //List cacheable for a day. More work to come on this, Etag, etc
	//fmt.Print("ListPhones")
	//m, _ := findAllPhones()

	phones := make([]map[string]interface{}, 0, 0)

	for _, v := range phonesCache {
		phones = append(phones, v)
	}

	return phones
}

func (serv PhoneService) GetPhone(id string) map[string]interface{} {
	//m := findAllPhones()
	return phonesCache[id]
}

var session *mgo.Session
var db *mgo.Database

var phonesCache map[string]map[string]interface{}

func main() {

	session, err := mgo.Dial("test:test@localhost")

	if err != nil {
		panic(err)
	}

	db = session.DB("test")
	defer session.Close()

	colls, _ := db.CollectionNames()
	for i, c := range colls {
		fmt.Printf("\n%v = %v", i, c)
	}

	//db.

	start()
	//importMongo()
}

func importMongo() {
	d, _ := ioutil.ReadFile("app/phones/phones.json")
	//var phones []Phone
	var newPhones []map[string]interface{}
	json.Unmarshal(d, &newPhones)
	fmt.Printf("\n%v", len(newPhones))
	//fmt.Printf("\n%v", newPhones)

	mapOfMap := findAllPhones()
	//var oldPhones []map[string]interface{}

	for _, p := range newPhones {
		if str, ok := p["id"].(string); ok {

			d, _ := ioutil.ReadFile("app/phones/" + str + ".json")
			var details map[string]interface{}
			json.Unmarshal(d, &details)

			for k, v := range p {
				details[k] = v
			}

			if mapOfMap[str] != nil {
				fmt.Printf("Updating %v", str)
				old := mapOfMap[str]
				err := db.C("phonesV2").Update(old, details)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Printf("Inserting %v", str)
				err := db.C("phonesV2").Insert(details)
				if err != nil {
					panic(err)
				}

			}

			//fmt.Printf("\n phone = %v", p)
			//fmt.Printf("\n details phone = %v", details)

		}
	}
}

func findAllPhones() map[string]map[string]interface{} {
	//var result []Phone

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	mapOfMap := make(map[string]map[string]interface{})
	var result []map[string]interface{}

	err := db.C("phonesV2").Find(nil).All(&result)
	if err != nil {
		//panic(err)
		fmt.Printf("\nError in get phones due to %v", err)
	}

	for _, m := range result {
		if str, ok := m["id"].(string); ok {
			//fmt.Printf("\nstr = %v", str)
			mapOfMap[str] = m
		}
	}

	//fmt.Printf("\n%v", string(test))
	return mapOfMap
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

	/*
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
	*/

	phonesCache = findAllPhones()

	gorest.RegisterService(new(PhoneService))
	/*http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
		listHandler(w, r)
	})*/
	//http.Handle("/html/", http.FileServer(d))
	http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("app"))
		r.URL.Path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
		//fmt.Printf("\nr.URL.Path=%v\n", r.URL.Path)
		fs.ServeHTTP(w, r)
	})
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(Address, nil)
}

/*
func listHandler(w http.ResponseWriter, r *http.Request) {

	file := strings.TrimPrefix(r.URL.Path, "/html/")
	//fmt.Printf("\nfile=%v", file)
	f, err := ioutil.ReadFile("app/" + file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(f)
}
*/
