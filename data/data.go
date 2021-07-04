package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type OsrmResponse struct{
	Code string `json:"code"`
	Routes []OsrmRoute `json:"routes"`

}
type OsrmRoute struct{
	Duration float64 `json:"duration"`
	Distance float64 `json:"distance"`


}
type Response struct{
	Source string
	Routes []Routes

}

type Routes struct{
	Duration float64
	Distance float64
	Destination string
}



func Source(w http.ResponseWriter, r *http.Request) string {
	src, ok := r.URL.Query()["src"]

	if !ok || len(src[0]) < 1 {
		err := "URLParam 'key' is missing"
		return err
	}

	res := src[0]


	return res



}

func Destination(w http.ResponseWriter, r *http.Request) []string{

	dst, ok := r.URL.Query()["dst"]

	if !ok || len(dst[0]) <1 {
		log.Fatalln("URLParam 'key' is missing")
		return nil
	}

	fmt.Println(dst)




	return dst
}

func Route(w http.ResponseWriter, r *http.Request) {
	alldst := Destination(w, r)
	src := Source(w, r)

	var endResponse Response
	endResponse.Source = src
	for _, dst := range alldst {

		req, err := http.NewRequest("GET", "http://router.project-osrm.org/route/v1/driving", nil)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}



		response, err := http.Get("http://router.project-osrm.org/route/v1/driving/" + src + ";" + dst)

		fmt.Println(req.URL.String())
		if err != nil {
			fmt.Printf("The HTTP requet failed with error %s\n", err)

		} else {
			var info Response
			data, _ := ioutil.ReadAll(response.Body)
			err := json.Unmarshal(data, &info)
			if err != nil {
				panic(err)
			}



			rdy2, err := json.Marshal(info)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(rdy2))
			var temp Routes
			temp.Duration = info.Routes[0].Duration
			temp.Distance = info.Routes[0].Distance
			temp.Destination = dst
			endResponse.Routes = append(endResponse.Routes, temp)





		}

	}
	sort.SliceStable(endResponse.Routes, func(i int, j int) bool {
		return (endResponse.Routes[i].Duration != endResponse.Routes[j].Duration &&
			endResponse.Routes[i].Duration < endResponse.Routes[j].Duration) ||
			(endResponse.Routes[i].Duration == endResponse.Routes[j].Duration &&
			endResponse.Routes[i].Distance < endResponse.Routes[j].Distance)

	})
	rdy, err := json.Marshal(endResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(rdy))
	fmt.Fprintf(w, string(rdy))



}

