package main

import "math"

const (
	EMPTYSTRING         string = ""
	ERROR                      = math.MaxFloat64 // larget float so that errors are sorted to last itmes in the list...
	MAXBUFFER                  = 1000
	MINBUFFER                  = 0
	CACHEPRUNEFREQUENCY        = 10
	CACHEEXPIRATION            = "60S"
)
