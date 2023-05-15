package main

type Extracted []struct {
	Destination latlong `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

type Response struct {
	Source latlong   `json:"source"`
	Routes Extracted `json:"routes"`
}
