// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestRouterWrapper(t *testing.T) {
	var routeSucceed bool
	wrapper := newWrapper(httprouter.New())

	handler := func(w http.ResponseWriter, r *http.Request) {
		routeSucceed = true
	}

	wrapper.GET("/", http.HandlerFunc(handler))
	wrapper.POST("/", http.HandlerFunc(handler))
	wrapper.PUT("/", http.HandlerFunc(handler))
	wrapper.PATCH("/", http.HandlerFunc(handler))
	wrapper.DELETE("/", http.HandlerFunc(handler))

	var httpVerbs = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}

	for _, method := range httpVerbs {
		routeSucceed = false
		req, err := http.NewRequest(method, "/", nil)

		if err != nil {
			t.Fatal(err)
			return
		}

		recorder := httptest.NewRecorder()

		wrapper.ServeHTTP(recorder, req)

		if routeSucceed != true {
			t.Fatalf("Expected wrapper method %s to succeed; it failed", method)
		}
	}

}
