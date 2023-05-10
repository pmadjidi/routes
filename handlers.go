package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("error"))
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("error"))
}

func (s *system) serve(w http.ResponseWriter, r *http.Request) {
	rawUri := r.RequestURI
	parsedUrl, err := url.Parse(rawUri)
	if err != nil {
		log.Println("Error parsing URL:", err)
		badRequest(w)
		return
	}
	queryValues := parsedUrl.Query()
	src := queryValues.Get("src")
	dsts, ok := queryValues["dst"]
	if !ok {
		badRequest(w)
		return
	}
	src, uDst, ok := validateSrcDsts(src, dsts)
	if !ok {
		badRequest(w)
		return
	}
	resp, err := s.callApi(src, uDst)
	if err != nil {
		internalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
