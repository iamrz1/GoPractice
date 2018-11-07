package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	//serve a string
	//fmt.Fprintf(w, "hello")

	//serve a string with extracted request
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	//serve a HTML file
	http.ServeFile(w, r, r.URL.Path[1:])
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	//convert value of i to string literal
	// and send to response writer
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
