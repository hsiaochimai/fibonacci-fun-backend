package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)



// Index : prints welcome
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello : prints hello name
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

//getFib produces a slice of integars based on the params given
func getFib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s, err := fibonacci(ps.ByName("number"))
	json.NewEncoder(w).Encode(s)
	if err != nil {
		fmt.Fprintf(w, "There is an error, %s!\n", err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/api", Index)
	router.GET("/api/hello/:name", Hello)
	router.GET("/api/fibonacci/:number", getFib)
	log.Fatal(http.ListenAndServe(":8080", router))
}
