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

func (s *system) callApiConcurrent(src latlong, dsts []latlong) (*Response, error) {
	requests := make([]*ApiPayload, 0)
	bodies := make([][]byte, 0)
	askApiForTheseDsts := make([]latlong, 0)

	// get cached bodies here...

	if s.enableCache {
		for _, dst := range dsts {
			cacheReq := s.newCacheRequest(src, dst)
			s.cacheRequest[cacheReq.bucket] <- cacheReq
			resp := <-cacheReq.resp
			if resp != nil {
				bodies = append(bodies, resp)
			} else {
				askApiForTheseDsts = append(askApiForTheseDsts, dst)
			}
		}
	} else {
		askApiForTheseDsts = dsts
	}
	for _, dst := range askApiForTheseDsts {
		ap := newApiPayload(src, dst)
		s.apiRequest <- ap
		requests = append(requests, ap)
	}

	for _, request := range requests {
		resp := <-request.resp
		if !resp.inError {
			bodies = append(bodies, resp.body)
		} else {
			bodies = append(bodies, []byte(EMPTYSTRING))
		}
	}

	resp, err := assmebleResponseFromBodies(bodies, src, dsts)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func assmebleResponseFromBodies(bodies [][]byte, src latlong, dsts []latlong) (*Response, error) {
	var resp Response
	var extracted = make(Extracted, len(dsts))
	if len(bodies) != len(dsts) {
		return nil, errors.New("api data not complete")
	}
	for i, body := range bodies {
		var apiData Api
		reader := bytes.NewReader(body)
		err := json.NewDecoder(reader).Decode(&apiData)

		if err != nil {
			log.Println("assmebleResponseFromBodies: ", err.Error())
			extracted[i].Destination = dsts[i]
			extracted[i].Duration = ERROR
			extracted[i].Distance = ERROR
		} else {
			extracted[i].Destination = dsts[i]
			extracted[i].Duration = apiData.Routes[0].Duration
			extracted[i].Distance = apiData.Routes[0].Distance
		}
	}
	sort.SliceStable(extracted, func(i, j int) bool {
		if extracted[i].Duration == extracted[j].Duration {
			return extracted[i].Distance < extracted[j].Distance
		}
		return extracted[i].Duration < extracted[j].Duration
	})

	resp.Source = src
	resp.Routes = extracted
	return &resp, nil
}
