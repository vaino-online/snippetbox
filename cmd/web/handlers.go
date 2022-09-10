package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// HTTP Handler functions

// index is a catch-all handler routed at "/"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the request URL path exactly matches "/". If it doesn't,
	// use the http.NotFound() function to send a 404 response to the
	// client.
	// Importantly, we then return from the handler, otherwise we would
	// also write the hello message.
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containing our templates' file paths
	// NOTE: Base template must be the first element.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Load the templates or freak out and throw a 500.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Render the base template.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id param from the query string and
	// try to convert it to an integer. If it fails, send a 404.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	// Interpolate the wanted snippet id into the response
	fmt.Fprintf(w, "Display snippet #%d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != http.MethodPost {
		// Add an "Allow: POST" header to the response header map.
		// This must be called before w.WriteHeader() or w.Write().
		w.Header().Set("Allow", http.MethodPost)
		// Send a Method Not Allowed HTTP error.
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet ..."))
}
