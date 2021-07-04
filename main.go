package main

import (
	"github.com/pawel/OSRM_API/data"
	"net/http"
)





func main() {
	http.HandleFunc("/routes", data.OsrmRouteCalculation)
	http.ListenAndServe(":8080", nil)
}
