package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
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
		saveData(counter)
	}
}

func saveData(counter Counter) {
	colQuerier := bson.M{"_id": counter.ID}
	change := bson.M{"$set": bson.M{"count": counter.Count}}
	err := col.Update(colQuerier, change)
	if err != nil {
		fmt.Println("error save data", err.Error())
		return
	}
}

func prepareConnection(url, username, pwd, db string) *mgo.Session {
	fmt.Println("creating new session with url", url)
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

func prepareConnectionWithUri(uri string) (ses *mgo.Session) {
	fmt.Println("creating new session with uri", uri)
	ci, err := mgo.ParseURL(uri)
	if err != nil {
		fmt.Println("failed to parse uri into dial info")
		return ses
	}
	fmt.Printf("%+v\n", ci)
	ses, err = mgo.DialWithInfo(ci)
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

	if os.Getenv("host") != "" {
		host = os.Getenv("host")
	}
	if os.Getenv("port") != "" {
		portDB = os.Getenv("port")
	}
	if os.Getenv("user") != "" {
		username = os.Getenv("user")
	}
	if os.Getenv("pwd") != "" {
		pwd = os.Getenv("pwd")
	}
	if os.Getenv("db") != "" {
		database = os.Getenv("db")
	}
	if os.Getenv("col") != "" {
		collname = os.Getenv("col")
	}

	url := host + ":" + portDB
	if os.Getenv("uri") != "" {
		url = os.Getenv("uri")
		conn = prepareConnectionWithUri(url)
	} else {
		conn = prepareConnection(url, username, pwd, database)
	}
	defer conn.Close()

	conn.SetMode(mgo.Monotonic, true)
	db := conn.DB(database)
	colList, err := db.CollectionNames()
	if err != nil {
		fmt.Println("error get collection list", err.Error())
	}
	isColExist := false
	for _, val := range colList {
		if val == collname {
			isColExist = true
			break
		}
	}
	col = conn.DB(database).C(collname)
	if !isColExist {
		counter := Counter{
			"mycounter",
			0,
		}
		saveData(counter)
	}

	http.HandleFunc("/", handlerFunc)
	port := "8003"
	fmt.Println("APPS run on port", port)
	fmt.Println("DATABASE run at host", url)
	http.ListenAndServe(":"+port, nil)
}
