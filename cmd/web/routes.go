package main

import "net/http"

func (app *application) routes() *http.ServeMux {
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

	return mux
}
