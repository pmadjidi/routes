package main

import (
	"strconv"
	"strings"
)

//Unsure of float is enough or float64 should be used

func validateLatLong(lt latlong) bool {
	coordinates := strings.Split(string(lt), ",")
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

func validateSrcDsts(src string, dsts []string) (latlong, []latlong, bool) {
	udsts := uniqueDsts(dsts)
	for _, dst := range udsts {
		if !validateLatLong(dst) {
			return latlong(EMPTYSTRING), nil, false
		}
	}

	if !validateLatLong(latlong(src)) {
		return latlong(EMPTYSTRING), nil, false
	}
	return latlong(src), udsts, true
}
