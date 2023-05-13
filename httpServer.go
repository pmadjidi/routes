package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func (s *system) startHttpOld() {
	err := http.ListenAndServe(":"+s.appPort, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("http server is closed...\n")
	} else if err != nil {
		fmt.Printf("failed to start the service... : %s\n", err)
		os.Exit(1)
	}
}

func (s *system) startHttp() {
	err := s.hServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("http server is closed...\n")
	} else if err != nil {
		fmt.Printf("failed to start the service... : %s\n", err)
		os.Exit(1)
	}
}
