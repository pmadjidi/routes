package main

import (
	"strconv"
	"strings"
)

//Unsure of float is enough or float64 should be used

func validateLatLong(latlong string) bool {
	coordinates := strings.Split(latlong, ",")
	if len(coordinates) != 2 {
		return false
	}

	if lat, err := strconv.ParseFloat(coordinates[0], 32); err != nil || (lat < -90.0 || lat > 90.0) {
		return false
	}

	if long, err := strconv.ParseFloat(coordinates[1], 32); err != nil || (long < -180.0 || long > 180.0) {
		return false
	}
	return true
}

func validateSrcDsts(src string, dsts []string) (string, []string, bool) {
	udsts := uniqueDsts(dsts)
	for _, dst := range udsts {
		if !validateLatLong(dst) {
			return EMPTYSTRING, nil, false
		}
	}
	if !validateLatLong(src) {
		return EMPTYSTRING, nil, false
	}
	return src, udsts, true
}
