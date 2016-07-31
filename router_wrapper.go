// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routerWrapper
// wrap julien's router http methods
type routerWrapper struct {
	router *httprouter.Router
}

// newRouterWrapper
// return pointer of routerWrapper
func newWrapper(r *httprouter.Router) *routerWrapper {
	return &routerWrapper{r}
}

// HTTP methods wrappers
func (rw routerWrapper) GET(path string, h http.Handler) {
	rw.router.GET(path, rw.wrapHandler(h))
}

func (rw routerWrapper) POST(path string, h http.Handler) {
	rw.router.POST(path, rw.wrapHandler(h))
}

func (rw routerWrapper) PUT(path string, h http.Handler) {
	rw.router.PUT(path, rw.wrapHandler(h))
}

func (rw routerWrapper) PATCH(path string, h http.Handler) {
	rw.router.PATCH(path, rw.wrapHandler(h))
}

func (rw routerWrapper) DELETE(path string, h http.Handler) {
	rw.router.DELETE(path, rw.wrapHandler(h))
}

// Wrap julien's httprouter.Handle inject params to a context
func (rw routerWrapper) wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	}
}

func (rw *routerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw.router.ServeHTTP(w, r)
}
