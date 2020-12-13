package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, This is the First Task")
}
func print(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) == 0 {
		fmt.Fprintf(w, "Url Params are missing ")
	} else{
		for k, v := range r.URL.Query() {
			fmt.Fprintf(w, "%s: %s \n", k, v)
		}
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/app", getInfo).Methods("GET")
	router.PathPrefix("/").HandlerFunc(print)
	http.Handle("/", router)
	http.ListenAndServe(":8083", router)
}
