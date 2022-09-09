package main

import (
	"log"
	"net/http"
)

// Configuration

// The address to bind the snippetbox srver to listen to.
// Generally you don't need to specify a host in the address unless your
// compter has multiple network interfaces and you just want to listen
// on one of them.
// NOTE: Should follow format "host:port".
const BindAddress = ":4000"

// HTTP Handler functions

// index is a catch-all handler routed at "/"
func index(w http.ResponseWriter, r *http.Request) {
	// Check if the request URL path exactly matches "/". If it doesn't,
	// use the http.NotFound() function to send a 404 response to the
	// client.
	// Importantly, we then return from the handler, otherwise we would
	// also write the hello message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet ..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet ..."))
}

// Main application

func main() {
	// Initialize a new servemux (aka router), then register the home
	// function as the handler for the index page
	mux := http.NewServeMux()

	// Go's servemux supports two different types of URL patterns: fixed
	// paths and subtree paths. Fixed paths don't end in /, subtree paths
	// do.
	// Subtree path patterns are matched (and handled) whenever the start
	// of a request URL matches the subtree path. Imagine a wildcard at
	// the end of them: / = /**, /static/ = /static/**.

	mux.HandleFunc("/", index)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/new", snippetCreate)

	// STart a new web server with the given network address to listen
	// on and the servemux we just created. If http.ListenAndServe()
	// returns an error, we use the log.Fatal() function to log the
	// error message and exit. Note that any error returned by
	// http.ListenAndServe() is always non-nil.
	log.Println("Snippetbox server listening on " + BindAddress)
	err := http.ListenAndServe(BindAddress, mux)
	log.Fatal(err)
}
