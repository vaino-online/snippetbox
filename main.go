package main

import (
	"log"
	"net/http"
)

// The address to bind the snippetbox srver to listen to.
// Generally you don't need to specify a host in the address unless your
// compter has multiple network interfaces and you just want to listen
// on one of them.
// NOTE: Should follow format "host:port".
const BindAddress = ":4000"

// home writes a byte slice containing a greeting as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Initialize a new servemux (aka router), then register the home
	// function as the handler for the index page
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// STart a new web server with the given network address to listen
	// on and the servemux we just created. If http.ListenAndServe()
	// returns an error, we use the log.Fatal() function to log the
	// error message and exit. Note that any error returned by
	// http.ListenAndServe() is always non-nil.
	log.Println("Snippetbox server listening on " + BindAddress)
	err := http.ListenAndServe(BindAddress, mux)
	log.Fatal(err)
}
