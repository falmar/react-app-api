// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandlerSuccess(t *testing.T) {
	jsonBody := []byte(`{"username":"http","password":"test"}`)

	// create request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	// recorder
	recorder := httptest.NewRecorder()

	// cast http.Handler to loginHandler
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(recorder, req)

	// check code
	if status := recorder.Code; status != http.StatusOK {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusOK, status)
	}
}

func TestLoginHandlerEmptyBody(t *testing.T) {
	// create request
	req, err := http.NewRequest("POST", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	// recorder
	recorder := httptest.NewRecorder()

	// cast http.Handler to loginHandler
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(recorder, req)

	// check code
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusBadRequest, status)
	}
}

func TestLoginHandlerBadBody(t *testing.T) {
	jsonBody := []byte("{bad-json}")

	// create request
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	// recorder
	recorder := httptest.NewRecorder()

	// cast http.Handler to loginHandler
	handler := http.HandlerFunc(loginHandler)

	handler.ServeHTTP(recorder, req)

	// check code
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusBadRequest, status)
	}
}
