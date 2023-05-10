package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

type system struct {
	apiUrl        string
	appPort       string
	serviceUrl    string
	optimizeLevel int
	timeOut       time.Duration
}

func createSys() system {
	timeout := os.Getenv("TIMEOUT")
	TIMEOUT, err := time.ParseDuration(timeout)
	if err != nil {
		TIMEOUT = 10 * time.Second
	}

	API_URL := os.Getenv("API_URL")
	if API_URL == EMPTYSTRING {
		log.Fatal("env API_URL not set... ")
	}

	PORT := os.Getenv("PORT")
	if PORT == EMPTYSTRING {
		log.Fatal("env PORT not set... ")
	}

	SERVICE_URL := os.Getenv("SERVICE_URL")
	if SERVICE_URL == EMPTYSTRING {
		log.Fatal("env SERVICE_URL not set... ")
	}

	optimizeLevel := os.Getenv("OPTIMIZE_LEVEL")
	OPTIMIZE_LEVEL, err := strconv.Atoi(optimizeLevel)
	if err != nil {
		OPTIMIZE_LEVEL = MINOPTLEVEL
	}

	if OPTIMIZE_LEVEL > MAXOPTLEVEL {
		OPTIMIZE_LEVEL = MAXOPTLEVEL
		log.Printf("optimization level above (%d) setting now to (%d)", MAXOPTLEVEL, MINOPTLEVEL)

	}

	distInMeters, err := precisionForLevel(OPTIMIZE_LEVEL)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Precision is set to level %d", OPTIMIZE_LEVEL)
	log.Printf("Clustring external API calls within distance %f (in meters)", distInMeters)

	return system{
		API_URL,
		PORT,
		SERVICE_URL,
		OPTIMIZE_LEVEL,
		TIMEOUT,
	}
}
