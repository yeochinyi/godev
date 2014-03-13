package main

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

type Phone struct {
	Age      int32
	Id       int32
	ImageUrl string
	Name     string
	Snippet  string
}

func main344() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	/*
		c := session.DB("test").C("people")
		err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			panic(err)
		}
	*/

	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	c := session.DB("test").C("phones")
	//var result []map[string]string
	var result []Phone

	err = c.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	test, _ := json.Marshal(result)
	fmt.Printf("Phones: %v", string(test))

	/*
		for i, v := range result {
			//fmt.Println("Phones:", v)
			test, _ := json.Marshal(v)
			fmt.Printf("\nPhones: %v", string(test))
			if i == 5 {
				break
			}
		}
	*/

	//fmt.Println("Phone:", result.Phone)
}
