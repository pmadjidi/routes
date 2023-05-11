package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"sort"
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

func (s *system) callApi(src string, dsts []string) (*Response, error) {
	bodies := make([][]byte, 0)
	for _, d := range dsts {

		// do not like to do string concatinations but this api looks not standard...
		url := s.apiUrl + src + ";" + d + "?overview=false"
		log.Println("calling: ", url)
		thisBody, err := s.apiCallHttp(url)
		if err != nil {
			return nil, err
		}
		bodies = append(bodies, thisBody)
	}
	resp, err := assmebleResponseFromBodies(bodies, src, dsts)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func assmebleResponseFromBodies(bodies [][]byte, src string, dsts []string) (*Response, error) {
	var resp Response
	var extracted = make(Extracted, len(dsts))
	if len(bodies) != len(dsts) {
		return nil, errors.New("Api data not complete")
	}
	for i, body := range bodies {
		var apiData Api
		reader := bytes.NewReader(body)
		err := json.NewDecoder(reader).Decode(&apiData)

		// all or nothing fail on partial info.... not sure how to handle partial fail...
		if err != nil {
			log.Println("assmebleResponseFromBodies: ", err.Error())
			return nil, err
		}
		extracted[i].Destination = dsts[i]
		extracted[i].Duration = apiData.Routes[0].Duration
		extracted[i].Distance = apiData.Routes[0].Distance
	}

	sort.SliceStable(extracted, func(i, j int) bool {
		return extracted[i].Duration < extracted[j].Duration
	})
	resp.Source = src
	resp.Routes = extracted
	return &resp, nil
}
