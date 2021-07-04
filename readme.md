OSRM API for Golang


A simple Go API for calculating distances and duration of travelling via car
between start and destination coordinates.

Using:
For proper use of this API you need to create a request similar to this example:
"GET http://localhost:8080/routes?src=,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219"


src=latitude(float64),longitude(float64)
dst=latitude(float64),longitude(float64)

You can set multiple destinations, but you always need to specify longitude and latitude.