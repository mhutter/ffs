# FallbackFileServer

[![GoDoc](https://godoc.org/github.com/mhutter/ffs?status.svg)](https://gowalker.org/github.com/mhutter/ffs)

HTTP handler that serves static files from a given directory, but serves `index.html` if the requested file is not found.

## Usage

Usage example using [mux][]

```go
r := mux.NewRouter()
r.PathPrefix("/api").Handler(api)

// Serve static assets
r.PathPrefix("/").
  Handler(ffs.New(publicDir)).
  Methods("GET")
```

[mux]: http://gorillatoolkit.org/pkg/mux

> [Manuel Hutter](https://hutter.io) -
> GitHub [@mhutter](https://github.com/mhutter) -
> Twitter [@dratir](https://twitter.com/dratir)
