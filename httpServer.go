package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func (s *system) startHttp() {
	http.HandleFunc(s.serviceUrl, s.serve)
	err := http.ListenAndServe(":"+s.appPort, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("http server is closed...\n")
	} else if err != nil {
		fmt.Printf("failed to start the service... : %s\n", err)
		os.Exit(1)
	}
}
