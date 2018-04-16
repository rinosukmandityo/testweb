package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var conn *mgo.Session
var col *mgo.Collection

type Counter struct {
	ID    string `json:"id" bson:"_id"`
	Count int
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		counter := Counter{}
		err := col.Find(nil).One(&counter)
		if err != nil {
			fmt.Println("error fetch data", err.Error())
			return
		}

		counter.Count++
		fmt.Fprint(w, fmt.Sprintf("<h1>Hello Container World! I have been seen %d times</h1>", counter.Count))
		fmt.Println(counter.ID, counter.Count)
		colQuerier := bson.M{"_id": counter.ID}
		change := bson.M{"$set": bson.M{"count": counter.Count}}
		err = col.Update(colQuerier, change)
		if err != nil {
			fmt.Println("error update data", err.Error())
			return
		}
	}
}

func prepareConnection() *mgo.Session {
	ci := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "local",
		Username: "",
		Password: "",
	}

	ses, err := mgo.DialWithInfo(ci)
	if err != nil {
		fmt.Println("failed to create mongo session")
		return ses
	}
	return ses
}

func main() {
	conn = prepareConnection()
	defer conn.Close()

	conn.SetMode(mgo.Monotonic, true)
	col = conn.DB("local").C("counter")

	http.HandleFunc("/", handlerFunc)
	port := "8003"
	fmt.Println("this server run at port", port)
	http.ListenAndServe(":"+port, nil)
}
