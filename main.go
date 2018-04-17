package main

import (
	"flag"
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

		colQuerier := bson.M{"_id": counter.ID}
		change := bson.M{"$set": bson.M{"count": counter.Count}}
		err = col.Update(colQuerier, change)
		if err != nil {
			fmt.Println("error update data", err.Error())
			return
		}
	}
}

func prepareConnection(url, username, pwd, db string) *mgo.Session {
	fmt.Println("creating new session")
	ci := &mgo.DialInfo{
		Addrs:    []string{url},
		Database: db,
		Username: username,
		Password: pwd,
	}

	ses, err := mgo.DialWithInfo(ci)
	if err != nil {
		fmt.Println("failed to create mongo session")
		return ses
	}
	fmt.Println("session created")
	return ses
}

func main() {
	host := "localhost"
	portDB := "27017"
	username := ""
	pwd := ""
	database := "local"
	collname := "counter"

	flag.StringVar(&host, "host", "localhost", "to determine database host")
	flag.StringVar(&portDB, "port", "27017", "to determine database port")
	flag.StringVar(&username, "user", "", "to determine database username")
	flag.StringVar(&pwd, "pwd", "", "to determine database password")
	flag.StringVar(&database, "db", "local", "to determine database name")
	flag.StringVar(&collname, "col", "counter", "to determine collection name")
	flag.Parse()

	url := host + ":" + portDB
	fmt.Println("url", url)
	conn = prepareConnection(url, username, pwd, database)
	defer conn.Close()

	conn.SetMode(mgo.Monotonic, true)
	col = conn.DB(database).C(collname)

	http.HandleFunc("/", handlerFunc)
	port := "8003"
	fmt.Println("APPS run on port", port)
	fmt.Println("DATABASE run at host", url)
	http.ListenAndServe(":"+port, nil)
}
