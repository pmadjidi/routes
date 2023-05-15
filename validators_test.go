package main

import "testing"

func TestValidateLatLong(t *testing.T) {
	var src latlong = "13.388860,52.517037"
	if !validateLatLong(src) {
		t.Fail()
	}
	var faultyLat latlong = "113.388860,52.517037"
	if validateLatLong(faultyLat) {
		t.Fail()
	}
	var faultyLong latlong = "13.388860,252.517037"
	if validateLatLong(faultyLong) {
		t.Fail()
	}
}

func TestValidateSrcDstsFault(t *testing.T) {
	src := "13.388860,52.517037"
	faultyLat := []string{"113.388860,52.517037", "13.388860,52.517037"}
	_, _, ok := validateSrcDsts(src, faultyLat)
	if ok {
		t.Fail()
	}

}

func TestValidateSrcDsts(t *testing.T) {
	src := "13.388860,52.517037"
	faultyLat := []string{"13.388860,52.517037", "13.388860,52.517037", "13.388860,52.517038"}
	_, vUdsts, ok := validateSrcDsts(src, faultyLat)
	if !ok {
		t.Fail()
	}
	if len(vUdsts) != 2 {
		t.Fail()
	}
}
