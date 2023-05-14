package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type system struct {
	apiUrl     string
	appPort    string
	serviceUrl string
	timeOut    time.Duration
	procNumber int
	apiRequest chan *ApiPayload
	shutDown   chan struct{}
	hServer    *http.Server
}

func (s *system) createTerminationHandler() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-term:
			for i := 0; i < s.procNumber; i++ {
				s.shutDown <- struct{}{}
			}
			break
		}
		s.shutDownHttpServer()
		os.Exit(1)
	}()
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

	NPROCESSORS := runtime.NumCPU()

	log.Printf("Number of processors: %d\n", NPROCESSORS)

	sys := system{
		API_URL,
		PORT,
		SERVICE_URL,
		TIMEOUT,
		NPROCESSORS,
		make(chan *ApiPayload, NPROCESSORS),
		make(chan struct{}),
		nil,
	}

	mux := http.NewServeMux()
	mux.Handle(sys.serviceUrl, &sys)

	httpServer := http.Server{Addr: ":" + sys.appPort, Handler: mux, WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second}

	sys.hServer = &httpServer
	return sys
}

func (s *system) initWorkers() {
	for i := 0; i < s.procNumber; i++ {
		go s.httpWorkers(i)
	}
}
