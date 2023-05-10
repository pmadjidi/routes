# routes

Clone the repo and to build:
cd to routes directory
run "go build"
and "./routes" to start
curl 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'
to run unittests run "go test"

Branch optimize, minize external api call by clustering destinations by means of geohashing coordinates and clustering based on optimization level. Implemented for fun...



