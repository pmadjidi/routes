package main

import (
	"errors"
	"strconv"
	"strings"
)

func uniqueDsts(dsts []string) []string {
	dict := make(map[string]bool)
	uDsts := make([]string, 0)
	for _, dst := range dsts {
		_, ok := dict[dst]
		if !ok {
			uDsts = append(uDsts, dst)
			dict[dst] = true
		}
	}
	return uDsts
}

func precisionForLevel(level int) (float32, error) {
	var resp float32 = -1.0

	switch level {
	case 1:
		resp = 5003530
	case 2:
		resp = 625441
	case 3:
		resp = 123264
	case 4:
		resp = 19545
	case 5:
		resp = 3803
	case 6:
		resp = 610
	case 7:
		resp = 118
	case 8:
		resp = 3.71
	case 9:
		resp = 5003530
	case 10:
		resp = 0.6
	}

	if resp == -1 {
		return resp, errors.New("precison out of range...")
	}
	return resp, nil
}

func getLatLong(ll string) (float64, float64, error) {
	v := strings.Split(ll, ",")
	lat, err := strconv.ParseFloat(v[0], 32)
	if err != nil {
		return 0, 0, err
	}
	long, err := strconv.ParseFloat(v[1], 32)
	if err != nil {
		return 0, 0, err
	}
	return lat, long, nil
}
