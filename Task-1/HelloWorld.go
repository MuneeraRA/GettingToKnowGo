package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, This is the First Task We work on ")
}
func start(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		fmt.Fprintf(w, "%s: %s \n", k, v)
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/app", getInfo).Methods("GET")
	router.HandleFunc("/", start).Methods("GET")
	http.ListenAndServe(":8083", router)
}
