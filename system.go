package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

type latlong string

type system struct {
	apiUrl       string
	appPort      string
	serviceUrl   string
	timeOut      time.Duration
	procNumber   int
	apiRequest   chan *ApiPayload
	cacheRequest []chan *CacheRequest
	setCache     []chan *CacheRequest
	shutDown     chan struct{}
	hServer      *http.Server
	enableCache  bool
}

func (s *system) numberOfSubsystems() int {
	// once for webworkers, and once for cache processors...
	return 2 * s.procNumber
}

func (s *system) createTerminationHandler() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-term:
			for i := 0; i < s.numberOfSubsystems(); i++ {
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
	bufferRequet := os.Getenv("BUFFER_REQUEST")
	BUFFER_REQUEST, err := strconv.Atoi(bufferRequet)
	if err != nil {
		BUFFER_REQUEST = 0
	} else {
		BUFFER_REQUEST = min(BUFFER_REQUEST, MAXBUFFER)
		if BUFFER_REQUEST < MINBUFFER {
			BUFFER_REQUEST = MINBUFFER
		}
	}

	enableCache := os.Getenv("ENABLE_CACHE")
	ENABLE_CACHE, err := strconv.ParseBool(enableCache)
	if err != nil {
		ENABLE_CACHE = false
	} else {
		if ENABLE_CACHE {
			log.Println("Enabling caching for DSTs...")
		}
	}

	log.Printf("Setting number of buffered requests to: %d\n", BUFFER_REQUEST)
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
	log.Printf("Number of processors detected: %d\n", NPROCESSORS)
	sys := system{
		API_URL,
		PORT,
		SERVICE_URL,
		TIMEOUT,
		NPROCESSORS,
		make(chan *ApiPayload, NPROCESSORS+BUFFER_REQUEST),
		make([]chan *CacheRequest, 0),
		make([]chan *CacheRequest, 0),
		make(chan struct{}),
		nil,
		ENABLE_CACHE,
	}

	for i := 0; i < sys.procNumber; i++ {
		chreq := make(chan *CacheRequest, sys.procNumber)
		chset := make(chan *CacheRequest, sys.procNumber)
		sys.cacheRequest = append(sys.cacheRequest, chreq)
		sys.setCache = append(sys.setCache, chset)
	}

	mux := http.NewServeMux()
	mux.Handle(sys.serviceUrl, &sys)

	httpServer := http.Server{Addr: ":" + sys.appPort, Handler: mux, WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second}

	sys.hServer = &httpServer
	return sys
}

func (s *system) initSubSystems() {
	log.Printf("starting one worker for each processeor total (%d)....\n", s.procNumber)
	for i := 0; i < s.procNumber; i++ {
		go s.httpWorkers(i)
		if s.enableCache {
			go s.cacheProcessor(i)
		}
	}
}
