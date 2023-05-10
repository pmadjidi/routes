package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"sort"

	geohash "github.com/mmcloughlin/geohash"
)

//Auto generated...

type any interface{}

type Api struct {
	Code   string `json:"code"`
	Routes []struct {
		Legs []struct {
			Steps    []any   `json:"steps"`
			Summary  string  `json:"summary"`
			Weight   float32 `json:"weight"`
			Duration float64 `json:"duration"`
			Distance float64 `json:"distance"`
		} `json:"legs"`
		WeightName string  `json:"weight_name"`
		Weight     float32 `json:"weight"`
		Duration   float64 `json:"duration"`
		Distance   float64 `json:"distance"`
	} `json:"routes"`
	Waypoints []struct {
		Hint     string    `json:"hint"`
		Distance float64   `json:"distance"`
		Name     string    `json:"name"`
		Location []float64 `json:"location"`
	} `json:"waypoints"`
}

func (s *system) clusterOnCoordinates(dsts []string) map[string][]string {
	buckets := make(map[string][]string)
	for _, dst := range dsts {
		// Ignoring errors, validated at http endpoint...
		lat, long, _ := getLatLong(dst)
		ghash := geohash.Encode(lat, long)
		log.Println("ghash: ", ghash)
		bucketTag := ghash[0:s.optimizeLevel]
		bucket, ok := buckets[bucketTag]
		if !ok {
			bucket := make([]string, 0)
			bucket = append(bucket, dst)
			buckets[bucketTag] = bucket
		} else {
			bucket = append(bucket, dst)
			buckets[bucketTag] = bucket
		}
	}

	if len(buckets) == len(dsts) {
		log.Printf("Clustring is not possible at optlevel %d", s.optimizeLevel)
	}
	return buckets
}

func (s *system) callApi(src string, dsts []string) (*Response, error) {
	bodies := make([][]byte, 0)
	clusteredDsts := s.clusterOnCoordinates(dsts)
	for _, d := range clusteredDsts {

		// do not like to do string concatinations but this api looks not standard...
		//Pick element zeo on each bucket to call the API
		url := s.apiUrl + src + ";" + d[0] + "?overview=false"
		log.Println("calling: ", url)
		thisBody, err := s.apiCallHttp(url)
		if err != nil {
			return nil, err
		}
		bodies = append(bodies, thisBody)
	}
	resp, err := assmebleResponseFromBodies(bodies, src, clusteredDsts, len(dsts))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func assmebleResponseFromBodies(bodies [][]byte, src string, clusteredDsts map[string][]string, sizeDst int) (*Response, error) {
	var resp Response
	var e = make(Extracted, sizeDst)
	if len(bodies) != len(clusteredDsts) {
		return nil, errors.New("Api data not complete")
	}

	bodyIndex := 0
	extractedIndex := 0
	for bucket, dsts := range clusteredDsts {
		log.Println("assembling for bucket: ", bucket)
		var apiData Api
		reader := bytes.NewReader(bodies[bodyIndex])
		err := json.NewDecoder(reader).Decode(&apiData)
		// all or nothing fail on partial info.... not sure how to handle partial fail...
		if err != nil {
			log.Println("assmebleResponseFromBodies: ", err.Error())
			return nil, err
		}
		for _, dst := range dsts {
			log.Println("assembling for dst: ", dst)
			e[extractedIndex].Destination = dst
			e[extractedIndex].Duration = apiData.Routes[0].Duration
			e[extractedIndex].Distance = apiData.Routes[0].Distance
			extractedIndex += 1
		}
		bodyIndex += 1
	}

	sort.SliceStable(e, func(i, j int) bool {
		return e[i].Duration < e[j].Duration
	})
	resp.Source = src
	resp.Routes = e
	return &resp, nil
}
