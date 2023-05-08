package main

import (
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
	TIMEOUT := os.Getenv("TIMEOUT")
	timeout, err := time.ParseDuration(TIMEOUT)
	if err != nil {
		timeout = 10 * time.Second
	}
	return system{
		os.Getenv("API_URL"),
		os.Getenv("PORT"),
		os.Getenv("SERVICE_URL"),
		timeout,
	}
}
