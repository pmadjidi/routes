package main

type Route struct {
	src string
	dst string
}

func (r *Route) validateSrcDsts(src string, dsts string) bool {
	if !validateLatLong(r.dst) || !validateLatLong(r.src) {
		return false
	}
	return true
}
