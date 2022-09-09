# Snippetbox

A basic pastebin for the purpose of learning the Go programming language.§

```
$ go run .
```

# Routes

```
Method    Pattern              Handler          Action
ANY       /                    index            Display the home page
ANY       /snippet/view?id=1   snippetView      Display a specific snippet
POST      /snippet/new         snippetCreate    Create a new snippet     
```