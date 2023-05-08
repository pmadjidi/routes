package main

import "testing"

func TestValidateLatLong(t *testing.T) {
	src := "13.388860,52.517037"
	if !validateLatLong(src) {
		t.Fail()
	}
	faultyLat := "113.388860,52.517037"
	if validateLatLong(faultyLat) {
		t.Fail()
	}
	faultyLong := "13.388860,252.517037"
	if validateLatLong(faultyLong) {
		t.Fail()
	}
}
