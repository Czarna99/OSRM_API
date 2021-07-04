package main

import (
	"github.com/pawel/OSRM_API/data"
	"net/http"
)





func main() {
	http.HandleFunc("/routes", data.Route)
	http.ListenAndServe(":8080", nil)
}
