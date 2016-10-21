// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestCheckTokenSuccess(t *testing.T) {
	// new request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// generate new token
	claims := MyClaims{
		User: User{Username: "dlavieri"},
	}
	claims.IssuedAt = time.Now().Unix()
	token, _ := generateToken(claims, jwtKey)

	// set token query param
	params := url.Values{}
	params.Set("token", token)
	req.URL.RawQuery = params.Encode()

	// new recorder
	recorder := httptest.NewRecorder()

	// http handler
	handler := http.HandlerFunc(checkTokenHandler)
	handler.ServeHTTP(recorder, req)

	// check code
	if status := recorder.Code; status != http.StatusOK {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusOK, status)
	}

	// decode response body
	response := &LoginResponse{}
	jsonRequestDecode(recorder.Body, response)

	// check token is the same
	if response.Token != token {
		t.Fatal("Expected handler to return same token as request param")
	}

	// check claims are not empty
	if response.Claims.User.Username == "" {
		t.Fatal("Expected handler to return valid claims")
	}

}

func TestCheckTokenEmpty(t *testing.T) {
	// new request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// set token query param
	params := url.Values{}
	params.Set("token", "")
	req.URL.RawQuery = params.Encode()

	// new recorder
	recorder := httptest.NewRecorder()

	// http handler
	handler := http.HandlerFunc(checkTokenHandler)
	handler.ServeHTTP(recorder, req)

	// check code must be bad request since token is empty
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusBadRequest, status)
	}
}

func TestCheckTokenRefresh(t *testing.T) {
	// new request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// generate new token
	claims := MyClaims{
		User: User{Username: "dlavieri"},
	}
	claims.IssuedAt = time.Now().Add(time.Minute * -90).Unix()
	token, _ := generateToken(claims, jwtKey)

	// set token query param
	params := url.Values{}
	params.Set("token", token)
	req.URL.RawQuery = params.Encode()

	// new recorder
	recorder := httptest.NewRecorder()

	// http handler
	handler := http.HandlerFunc(checkTokenHandler)
	handler.ServeHTTP(recorder, req)

	// check code must be 200 Ok
	if status := recorder.Code; status != http.StatusOK {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusOK, status)
	}

	response := &LoginResponse{}
	jsonRequestDecode(recorder.Body, response)

	iat := time.Since(time.Unix(response.Claims.IssuedAt, 0))

	// token should have been recently generated
	if iat.Minutes() > 1 {
		t.Fatal("Expected handler to generate new token")
	}
}

func TestCheckTokenBadJSON(t *testing.T) {
	// new request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// set token query param
	params := url.Values{}
	params.Set("token", "{oops}")
	req.URL.RawQuery = params.Encode()

	// new recorder
	recorder := httptest.NewRecorder()

	// http handler
	handler := http.HandlerFunc(checkTokenHandler)
	handler.ServeHTTP(recorder, req)

	// check code must be StatusUnauthorized since token is bad formed
	if status := recorder.Code; status != http.StatusUnauthorized {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusUnauthorized, status)
	}
}

func TestCheckTokenBadClaims(t *testing.T) {
	// new request
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// generate new token
	claims := jwt.MapClaims{
		"thingy": "not really MyClaims",
	}

	token, _ := generateToken(claims, jwtKey)

	// set token query param
	params := url.Values{}
	params.Set("token", token)
	req.URL.RawQuery = params.Encode()

	// new recorder
	recorder := httptest.NewRecorder()

	// http handler
	handler := http.HandlerFunc(checkTokenHandler)
	handler.ServeHTTP(recorder, req)

	// check code must be StatusUnauthorized Ok
	if status := recorder.Code; status != http.StatusUnauthorized {
		t.Fatalf("Expected HTTP Status: %d; Got: %d", http.StatusUnauthorized, status)
	}
}
