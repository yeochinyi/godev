package json

import (
	"fmt"
	//	"html/template"
	//"io/ioutil"
	"net/http"
	//"os"
	//"regexp"
	"code.google.com/p/gorest"
	//"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"strings"
)

const (
	Address string = ":9090"
)

var session *mgo.Session
var db *mgo.Database

func Start() {

	session, err := mgo.Dial("test:test@localhost")

	if err != nil {
		panic(err)
	}

	db = session.DB("test")
	defer session.Close()

	gorest.RegisterService(new(CollectionService))

	http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("app"))
		r.URL.Path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
		fmt.Printf("\nr.URL.Path=%v\n", r.URL.Path)
		fs.ServeHTTP(w, r)
	})
	http.Handle("/rest", gorest.Handle())

	fmt.Printf("Status is %v", http.ListenAndServe(Address, nil))
}

type CollectionService struct {
	gorest.RestService `root:"/rest/" consumes:"application/json" produces:"application/json"`
	helloWorld         gorest.EndPoint `method:"GET" path:"/helloworld/" output:"string"`
	listCnames         gorest.EndPoint `method:"GET" path:"/cnames/" output:"[]string"`
	findC              gorest.EndPoint `method:"GET" path:"/c/{cname:String}/{id:String}" output:"[]string"`
}

func (serv CollectionService) HelloWorld() string {
	return "Hello World"
}

func (serv CollectionService) listCnames() []string {
	//serv.ResponseBuilder().CacheMaxAge(60*60*24) //Cacheable for a day
	l, _ := db.CollectionNames()
	for i := range l {
		fmt.Println(i)
	}

	return l
}

func (serv CollectionService) findC(cname String, id String) []string {
	var search bson.M
	if id == nil {
		search = bson.M{"_id": id}
	}
	var result []map[string]interface{}
	err := db.C(cname).Find(search).All(&result)
	if err != nil {
		fmt.Printf("\nError in get "+cname+" due to %v", err)
	}

	return l
}
