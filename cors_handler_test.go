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

// should allow all the httpVerbs described
func TestCorsOriginFullMethods(t *testing.T) {
	var headerKey string
	var expectedHeader string
	// omit GET
	var httpVerbs = []string{"POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	cors := CORS{nil}

	// create request
	req, err := http.NewRequest("OPTIONS", "/", nil)

	if err != nil {
		t.Fatal(err)
		return
	}

	for _, method := range httpVerbs {
		expectedHeader = method

		// Set PreFligt headers
		req.Header.Set("Origin", "http://example.com")
		req.Header.Set("Access-Control-Allow-Methods", method)

		// Response Recorder
		rr := httptest.NewRecorder()

		cors.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("Expected status %d; Got: %d", http.StatusOK, status)
		}

		responseHeader := rr.Header()
		headerKey = "Access-Control-Allow-Methods"

		if header := responseHeader.Get(headerKey); header != expectedHeader {
			t.Fatalf(
				"Expected Header %s Value: %s; Got: %s",
				headerKey,
				expectedHeader,
				header,
			)
		}

		headerKey = "Access-Control-Allow-Origin"
		expectedHeader = "http://example.com"
		if header := responseHeader.Get(headerKey); header != expectedHeader {
			t.Fatalf(
				"Expected Header %s Value: %s; Got: %s",
				headerKey,
				expectedHeader,
				header,
			)
		}

		headerKey = "Access-Control-Allow-Headers"
		expectedHeader = "Content-Type, Origin"
		if header := responseHeader.Get(headerKey); header != expectedHeader {
			t.Fatalf(
				"Expected Header %s Value: %s; Got: %s",
				headerKey,
				expectedHeader,
				header,
			)
		}
	}
}

// should not set any PreFligt CORS header since there is no Origin
func TestCorsWithNoOrigin(t *testing.T) {
	var routerSucceed bool

	router := httprouter.New()

	handler := func(_ http.ResponseWriter, _ *http.Request) {
		routerSucceed = true
	}

	router.HandlerFunc("GET", "/", handler)

	cors := CORS{router}

	// create request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
		return
	}

	// No PreFligt headers

	// Response Recorder
	rr := httptest.NewRecorder()

	cors.ServeHTTP(rr, req)

	responseHeader := rr.Header()

	headerKeys := []string{
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Headers",
	}

	for _, headerKey := range headerKeys {
		if header := responseHeader.Get(headerKey); header != "" {
			t.Fatalf("Expected header %s to be empty; Got: %s", headerKey, header)
		}
	}

	if !routerSucceed {
		t.Fatal("Expected httprouter to ServeHTTP to be executed")
	}
}
