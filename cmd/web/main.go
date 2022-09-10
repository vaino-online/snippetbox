package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold app-wide dependencies.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// Main server setup

func main() {
	// Parse configuration from command-line flags.
	addr := flag.String("addr", ":4000", "Server bind address")
	flag.Parse()

	// Create a logger for informational messages. This takes three args:
	// the defination to write the logs to (os.Stdout), a string prefix for
	// all log messages (INFO followed by \t), and flags to indicate what
	// additional information to include (local date and time (bitmask)).
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create another logger, but for errors, using os.Stderr and use
	// the log.Lshortfile to include the error location in a file.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize the application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initialize a new servemux (aka router) to store our URL paths.
	mux := http.NewServeMux()

	// Create a static file server. The file server sanitizes all request
	// paths by running them through path.Clean() first. It strips . and ..
	// to block directory traversal attacks.
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
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/new", app.snippetCreate)

	// Initialize a new http.Server struct with proper bind targets
	// and loggers.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Start a new web server with the given network address to listen
	// on and the servemux we just created. If http.ListenAndServe()
	// returns an error, we use the log.Fatal() function to log the
	// error message and exit. Note that any error returned by
	// http.ListenAndServe() is always non-nil.
	infoLog.Printf("Starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
