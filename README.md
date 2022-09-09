# Snippetbox

A basic pastebin for the purpose of learning the Go programming language.§

## Running

```
$ go run ./cmd/web
```

## Routes

```
Method    Pattern              Handler          Action
ANY       /                    index            Display the home page
ANY       /snippet/view?id=1   snippetView      Display a specific snippet
POST      /snippet/new         snippetCreate    Create a new snippet     
ANY       /static/             http.FileServer  Serve static files
```

## Structure

The `cmd` directory contains application-specific code for the executable applications in the project.

The `internal` directory contains ancillary non-application-specific code used in the project. Holds reusable helpers and SQL database models etc.

The `ui` directory contains user-interface assets used by the web app. Specifically, the `ui/html` directory will contain HTML templates, and the `ui/static` will contain static assets like images and stylesheets.