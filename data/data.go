package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)
//OsrmResponse - struct that saves response information received in JSON format from OSRM router
type OsrmResponse struct{
	Code string `json:"code"`
	Routes []OsrmRoute `json:"routes"`

}
//OsrmRoute - routes for OsrmResponse struct
type OsrmRoute struct{
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`


}
//Response - struct for displaying proper information received by OSRM router by  Go API
type Response struct{
	Source string
	Routes []Routes

}
//Routes - routes struct for Response
type Routes struct{
	Duration float64
	Distance float64
	Destination string
}


//GetDestinationFromUrl - function that receiving destination  longitude and latitude from API GET request,
//returning them as string.
func GetSourceFromUrl(w http.ResponseWriter, r *http.Request) string {
	src, ok := r.URL.Query()["src"]

	if !ok || len(src[0]) < 1 {
		err := "URLParam 'key' is missing"
		return err
	}

	source := src[0]


	return source



}

//GetSourceFromUrl - function that receiving source longitude and latitude from API GET request,
//returning them as slice.
func GetDestinationFromUrl(w http.ResponseWriter, r *http.Request) []string{

	destination, ok := r.URL.Query()["dst"]

	if !ok || len(destination[0]) <1 {
		log.Fatalln("URLParam 'key' is missing")
		return nil
	}


	return destination
}
//OsrmRouteCalculation - function that is using provided source and destination coordinates and sending them to OSRM router
func OsrmRouteCalculation(w http.ResponseWriter, r *http.Request) {
	allDestination := GetDestinationFromUrl(w, r)
	source := GetSourceFromUrl(w, r)

	var endResponse Response
	endResponse.Source = source
	for _, dst := range allDestination {

		response, err := http.Get("http://router.project-osrm.org/route/v1/driving/" + source + ";" + dst)

		if err != nil {
			fmt.Printf("The HTTP requet failed with error %s\n", err)

		} else {
			//Reading response body and unmarshalling it into info(type Response struct)
			var info Response
			data, _ := ioutil.ReadAll(response.Body)
			err := json.Unmarshal(data, &info)
			if err != nil {
				panic(err)
			}




			var temp Routes
			temp.Duration = info.Routes[0].Duration
			temp.Distance = info.Routes[0].Distance
			temp.Destination = dst
			endResponse.Routes = append(endResponse.Routes, temp)


		}

	}
	//Sorting from smaller to highest duration. If duration is same, data will be sorted via lower to higher distance
	sort.SliceStable(endResponse.Routes, func(i int, j int) bool {
		return (endResponse.Routes[i].Duration != endResponse.Routes[j].Duration &&
			endResponse.Routes[i].Duration < endResponse.Routes[j].Duration) ||
			(endResponse.Routes[i].Duration == endResponse.Routes[j].Duration &&
			endResponse.Routes[i].Distance < endResponse.Routes[j].Distance)

	})
	//Marshalling needed data received from OSRM router
	requestedData, err := json.Marshal(endResponse)
	if err != nil {
		panic(err)
	}
	//Displaying data into web browser
	fmt.Fprintf(w, string(requestedData))



}

