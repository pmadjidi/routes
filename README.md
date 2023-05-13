## Getting Started

To use this repository, follow the instructions below:

1. Clone the repository.
2. Navigate to the `routes` directory.
3. Run the command `go build` to build the project.
4. Start the application by running `./routes`.
5. Use curl to test the application: `curl 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'`.
6. To run unit tests, execute `go test`.

## Branch: optimize

The `optimize` branch aims to minimize external API calls by clustering destinations. This is achieved by geohashing coordinates and clustering based on the optimization level. The algorithm selects a certain length of the hash (precision) for bucketing destinations. Please note that this implementation is done purely for fun and experimentation purposes.

## Branch: concurrent

Branch "concurrent" is a multi-processor implementation for the "routes" service. The aim of this implementation is to handle external API calls in parallel as much as possible.
