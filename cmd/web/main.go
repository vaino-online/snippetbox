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

	// Initialize a new http.Server struct with proper bind targets
	// and loggers.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
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
