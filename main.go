package main

import (
	"fmt"
	"net/http"
)


func routes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
	}
	w.Write([]byte("Received a GET request\n"))
}


func main() {
	http.HandleFunc("/", routes)
	http.ListenAndServe(":8000", nil)
}
