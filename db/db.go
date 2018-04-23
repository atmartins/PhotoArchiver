package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

var db *mgo.Database
var session *mgo.Session

// Connect establishes a persistent connection to Mongo
func Connect(addr string, dbName string) (err error) {
	session, err = mgo.Dial(addr)
	if err != nil {
		return
	}
	// Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)
	db = session.DB(dbName)
	return
}

// Disconnect from Mongo
func Disconnect() {
	session.Close()
}

// Test insert
func Test() (err error) {
	c := db.C("people")
	err = c.Insert(
		&Person{"Aaron", "+541-515-4930"},
		&Person{"Cla", "+55 53 8402 8510"},
	)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Aaron"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
	return
}
