package main
import (
    "fmt"
"github.com/julienschmidt/httprouter"
 "net/http"
"log"

)
// Index : prints welcome
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) { 
	fmt.Fprint(w, "Welcome!\n")
}
// Hello : prints hello name 
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { 
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}





func main() {
router := httprouter.New() 
router.GET("/api", Index) 
router.GET("/api/hello/:name", Hello)
log.Fatal(http.ListenAndServe(":8080", router)) }