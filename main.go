package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, fmt.Sprintf("<h1>You access this page at %s</h1>", time.Now().Format("15:04:05")))
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	port := "8003"
	fmt.Println("this server run at port", port)
	http.ListenAndServe(":"+port, nil)
}
