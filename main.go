package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Routers...")
	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("Server started and listening on localhost port 9002")
	log.Fatal(http.ListenAndServe(":9002", nil))

}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is GET")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is POST")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is DELETE")
}
