package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Main: started application on port 9002!")
	http.HandleFunc("/", helloworld)
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world \n How are you doing today ?")
}
