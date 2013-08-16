package main

import (
	"log"
	"net/http"
	"time"
)

type WatchHandler struct{}

func (wc *WatchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.URL.Path)

	// Look for it in all the places.
	time.Sleep(5 * time.Second)
	// Didn't find anything.
	http.NotFound(w, req)
}
