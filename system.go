package main

import (
	"log"
	"os"
	"time"
)

type system struct {
	apiUrl     string
	appPort    string
	serviceUrl string
	timeOut    time.Duration
}

func createSys() system {
	timeout := os.Getenv("TIMEOUT")
	TIMEOUT, err := time.ParseDuration(timeout)
	if err != nil {
		TIMEOUT = 10 * time.Second
	}

	API_URL := os.Getenv("API_URL")
	PORT := os.Getenv("PORT")
	SERVICE_URL := os.Getenv("SERVICE_URL")

	if API_URL == EMPTYSTRING {
		log.Fatal("env API_URL not set... ")
	}

	if PORT == EMPTYSTRING {
		log.Fatal("env PORT not set... ")
	}

	if SERVICE_URL == EMPTYSTRING {
		log.Fatal("env SERVICE_URL not set... ")
	}

	return system{
		API_URL,
		PORT,
		SERVICE_URL,
		TIMEOUT,
	}
}
