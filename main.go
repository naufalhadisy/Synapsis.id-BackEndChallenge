package main

import (
	//"fmt"
	//"html/template"
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	//sw "./go"

	"github.com/gorilla/mux"
)

/*//var templates *template.Template

func main() {

	//templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/start", startHandler).Methods("GET")
	r.HandleFunc("/stop", stopHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func startHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Starting App")
	//templates.ExecuteTemplate(w, "index.html", nil)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Stopping App")
}*/

func main() {
	log.Printf("Server started")

	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
