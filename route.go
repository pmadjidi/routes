package main

type Route struct {
	src latlong
	dst latlong
}

func (r *Route) validateSrcDsts(src latlong, dsts latlong) bool {
	if !validateLatLong(r.dst) || !validateLatLong(r.src) {
		return false
	}
	return true
}
