package main

import (
	"fmt"
	"log"
	"net/http"

	"goapi/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("HELLO UNICORNS!")
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	log.Fatal(http.ListenAndServe(":9001", r))
}
