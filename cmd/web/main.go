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

// Main server setup

func main() {
	// Initialize a new servemux (aka router) to store our URL paths.
	mux := http.NewServeMux()

	// Create a static file server.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as the handler for all URL paths
	// that start with "/static/". Strip the prefix before passing
	// the request to the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Go's servemux supports two different types of URL patterns: fixed
	// paths and subtree paths. Fixed paths don't end in /, subtree paths
	// do.
	// Subtree path patterns are matched (and handled) whenever the start
	// of a request URL matches the subtree path. Imagine a wildcard at
	// the end of them: / = /**, /static/ = /static/**.
	// NOTES:
	// - Longer URL patterns take precedence over shorter ones.
	// - Request URL paths are automatically sanitized and the user
	//   redirected (if needed).
	// - If a subtree path is registered and a request to it without a
	//   trailing slash is received, user will be sent a 301 Permanent
	//   Redirect to the subtree path with the slash added.
	// - URL patterns can contain hostnames like this:
	//   mux.HandleFunc("foo.vaino.lol/", fooHandler)
	//   mux.HandleFunc("bar.vaino.lol/", barHandler)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/new", snippetCreate)

	// Start a new web server with the given network address to listen
	// on and the servemux we just created. If http.ListenAndServe()
	// returns an error, we use the log.Fatal() function to log the
	// error message and exit. Note that any error returned by
	// http.ListenAndServe() is always non-nil.
	log.Println("Snippetbox server listening on " + BindAddress)
	err := http.ListenAndServe(BindAddress, mux)
	log.Fatal(err)
}
