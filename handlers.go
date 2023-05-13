package main

import (
	"context"
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

func (s *system) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	resp, err := s.callApiConcurrent(src, uDst)
	if err != nil {
		internalError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *system) shutDownHttpServer() {
	log.Fatalln("terminating http server.....")
	err := s.hServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
}
