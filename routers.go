// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"

	"github.com/go-chi/chi/v5"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

// flag indicating if the normal or the test handler should be loaded
var loadTestHandler = false

func init() {
	runtime.GOMAXPROCS(1)

	log.SetOutput(new(mockResponseWriter))
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

func httpHandlerFuncTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

// chi
func chiHandleWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, chi.URLParam(r, "name"))
}

func loadChi(routes []route) http.Handler {
	h := httpHandlerFunc
	if loadTestHandler {
		h = httpHandlerFuncTest
	}

	re := regexp.MustCompile(":([^/]*)")

	mux := chi.NewRouter()
	for _, route := range routes {
		path := re.ReplaceAllString(route.path, "{$1}")

		switch route.method {
		case "GET":
			mux.Get(path, h)
		case "POST":
			mux.Post(path, h)
		case "PUT":
			mux.Put(path, h)
		case "PATCH":
			mux.Patch(path, h)
		case "DELETE":
			mux.Delete(path, h)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadChiSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := chi.NewRouter()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknown HTTP method: " + method)
	}
	return mux
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
