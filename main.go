package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirectToHttps(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	// Redirect the incoming HTTP request.
	http.Redirect(w, req, target, http.StatusMovedPermanently)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}

func main() {
	http.HandleFunc("/", handler)
	// Start the HTTP server and redirect all incoming connections to HTTPS
	if err := http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps)); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
