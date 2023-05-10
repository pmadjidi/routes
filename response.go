package main

type Extracted []struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

type Response struct {
	Source string    `json:"source"`
	Routes Extracted `json:"routes"`
}
