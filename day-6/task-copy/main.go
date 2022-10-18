package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	route.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("about"))
	})

	route.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("help me"))
	})

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}
