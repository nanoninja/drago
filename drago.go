// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package drago is an idiomatic HTTP Middleware for Go.
package drago

import "net/http"

// Middleware represents a piece of the chain.
//
//    func DemoMiddleware(next http.Handler) http.Handler {
//        return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
//            fmt.Println("Demo Middleware")
//            next.ServeHTTP(rw, r)
//        })
//    }
//
type Middleware func(http.Handler) http.Handler

// Chain is a stack of Middleware.
type Chain []Middleware

// New returns a new Chain instance with no middleware preconfigured.
//
//    handler := drago.New(Middleware1, Middleware2).Handler(http.NewServeMux())
//
func New(mw ...Middleware) Chain {
	c := make(Chain, 0)
	c.Use(mw...)
	return c
}

// Use adds a Middleware onto the chain stack.
// Middleware are invoked in the order they are added to the chain.
//
//    d := drago.New()
//    d.Use(Middleware1, Middleware2).Handler(handler)
//
func (c *Chain) Use(mw ...Middleware) {
	*c = append(*c, mw...)
}

// Extend allows the use to join two chains of middlewares.
func (c *Chain) Extend(chain Chain) {
	c.Use(chain...)
}

// Handler chains the middleware and returns a http.Handler to
// use for the given request.
func (c Chain) Handler(h http.Handler) http.Handler {
	if h == nil {
		h = http.NewServeMux()
	}
	for i := len(c) - 1; i >= 0; i-- {
		h = c[i](h)
	}
	return h
}

// HandlerFunc is an adapter to allow the use of
// ordinary functions as HTTP handlers.
func (c Chain) HandlerFunc(h http.HandlerFunc) http.Handler {
	if h == nil {
		return c.Handler(nil)
	}
	return c.Handler(h)
}
