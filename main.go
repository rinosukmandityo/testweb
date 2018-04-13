package main

import (
	"fmt"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"net/http"
)

var conn dbox.IConnection

type Counter struct {
	Count int
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		csr, err := conn.NewQuery().From("counter").Cursor(nil)
		if err != nil {
			fmt.Println("error create cursor", err.Error())
			return
		}
		counter := Counter{}
		err = csr.Fetch(&counter, 1, false)
		if err != nil {
			fmt.Println("error fetch data", err.Error())
			return
		}
		counter.Count++
		fmt.Fprint(w, fmt.Sprintf("<h1>Hello Container World! I have been seen %d times</h1>", counter.Count))
		err = conn.NewQuery().From("counter").Save().Exec(map[string]interface{}{
			"data": Counter{counter.Count},
		})
		if err != nil {
			fmt.Println("error save data", err.Error())
			return
		}
	}
}

func prepareConnection() {
	ci := &dbox.ConnectionInfo{"localhost:27017", "local", "", "", nil}
	var err error
	conn, err = dbox.NewConnection("mongo", ci)
	if err != nil {
		return
	}
	err = conn.Connect()
	if err != nil {
		return
	}
	return
}

func main() {
	prepareConnection()
	http.HandleFunc("/", handlerFunc)
	port := "8003"
	fmt.Println("this server run at port", port)
	http.ListenAndServe(":"+port, nil)
}
