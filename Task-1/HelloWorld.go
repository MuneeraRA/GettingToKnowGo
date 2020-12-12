package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/App", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've completed your First Task")
	})

	http.HandleFunc("/x", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've completed your First Task \n")
	})

	http.ListenAndServe(":80", nil)
}
