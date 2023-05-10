package main

import (
	"fmt"
	"testing"
)

func TestUniqueDsts(t *testing.T) {
	dsts := []string{"13.428555,52.523219", "13.428555,52.523219", "13.428555,52.523219"}
	uDsts := uniqueDsts(dsts)
	if len(uDsts) != 1 {
		fmt.Printf("%+v", uDsts)
		t.Fail()
	}
}
