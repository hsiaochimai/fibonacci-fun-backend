package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Fib struct (Model)
// type Fib struct {
// 	input string `json:"input"`
// 	value string `json:"value"`
// }

// Index : prints welcome
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello : prints hello name
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func getFib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(fibonacci(ps.ByName("number")))
}

func main() {
	router := httprouter.New()
	router.GET("/api", Index)
	router.GET("/api/hello/:name", Hello)
	router.GET("/api/fib/:number", getFib)
	log.Fatal(http.ListenAndServe(":8080", router))
}
