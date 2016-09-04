# Drago

Drago is an idiomatic HTTP Middleware for Go.

[![license](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/nanoninja/drago/blob/master/LICENSE) [![godoc](https://godoc.org/github.com/nanoninja/drago?status.svg)](https://godoc.org/github.com/nanoninja/drago) [![build status](https://travis-ci.org/nanoninja/drago.svg)](https://travis-ci.org/nanoninja/drago)  [![go report card](https://goreportcard.com/badge/github.com/nanoninja/drago)](https://goreportcard.com/report/github.com/nanoninja/drago) [![codebeat](https://codebeat.co/badges/58e89ce4-2fd8-4a93-b624-afdbbb44a6e3)](https://codebeat.co/projects/github-com-nanoninja-drago)

## Installation

    go get github.com/nanoninja/drago

## Getting Started

After installing Go and setting up your
[GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file.

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/nanoninja/drago"
)

func DemoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Demo Middleware Before")
		next.ServeHTTP(rw, r)
        fmt.Println("Demo Middleware After")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome to the hom page!")
	})

	handler := drago.New(DemoMiddleware).Handler(mux)

	http.ListenAndServe(":3000", handler)
}
```

## Usage examples

### Use
``` go
d := drago.New()

d.Use(Middleware1, Middleware2)

handler := d.Handler(http.NewServeMux())
```

### Extend
``` go
a := drago.New(Middleware1, Middleware2)
b := drago.New(Middleware3, Middleware4)

a.Extend(b) // a == Middleware1, Middleware2, Middleware3, Middleware4

handler := a.Handler(http.NewServeMux())
```

### HandlerFunc
``` go
d := drago.New(Middleware1, Middleware2)

handler := d.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(rw, "Welcome!")
})
```

## Third Party Middleware

Here is a current list of Drago compatible middleware :

| Middleware | Author | Description |
| -----------|--------|-------------|
| [Bulma](https://github.com/nanoninja/bulma) | [Vincent Letourneau](https://github.com/nanoninja) | Basic authentication implementation for Go. |

## License

Drago is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a [BSD license](https://github.com/nanoninja/drago/blob/master/LICENSE).
