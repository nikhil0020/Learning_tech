package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello from AM-SAPP")
	r := mux.NewRouter()

	r.HandleFunc("/", serveRoute).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Oneforuth.com</h1>"))
}
