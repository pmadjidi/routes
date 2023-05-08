# routes
test app 

clone the repo
cd to routes directory
go build 
create .env file in the routes directory with the content:
"
PORT=8080
API_URL= http://router.project-osrm.org/route/v1/driving/
SERVICE_URL = /routes
TIMEOUT = 10s
"

./routes to start

curl 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'

to run unittests run "go test"



