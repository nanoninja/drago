// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package drago

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v), Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestDragoUse(t *testing.T) {
	result := ""
	response := httptest.NewRecorder()

	d := New()
	d.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			result += "A"
			next.ServeHTTP(rw, r)
			result += "B"
		})
	})
	d.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			result += "C"
			next.ServeHTTP(rw, r)
			result += "D"
		})
	})
	d.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		result += "E"
		rw.WriteHeader(http.StatusInternalServerError)
	}).ServeHTTP(response, (*http.Request)(nil))

	expect(t, result, "ACEDB")
	expect(t, response.Code, http.StatusInternalServerError)
}

func TestDragoHandlerNil(t *testing.T) {
	result := ""
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	d := New()
	d.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			result += "Go"
			next.ServeHTTP(rw, r)
		})
	})
	d.HandlerFunc(nil).ServeHTTP(response, request)

	expect(t, result, "Go")
	expect(t, response.Code, http.StatusNotFound)
}

func TestDragoExtend(t *testing.T) {
	result := 0
	d := New()
	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			result += 2
			next.ServeHTTP(rw, r)
		})
	}
	expect(t, len(d), 0)
	d.Use(mw, mw)
	expect(t, len(d), 2)
	d.Extend(New(mw, mw))
	expect(t, len(d), 4)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	d.Handler(nil).ServeHTTP(response, request)
	expect(t, result, len(d)*2)
}
