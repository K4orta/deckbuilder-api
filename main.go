package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	router.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello Users")
	})

	router.HandleFunc("/decks", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello Decks")
	})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":9601")
}
