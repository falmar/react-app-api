// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"time"
)

// LoginResponse response given to request
type LoginResponse struct {
	Claims MyClaims `json:"claims"`
	Token  string   `json:"token"`
}

// LoginRequest incoming request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// set Content-Type
	w.Header().Set("Content-Type", "application/json")

	// objects to marshal/unmarshal json
	request := LoginRequest{}
	response := LoginResponse{}

	// TODO: replace this ugly closure with real error handler
	badRequest := func(message string) {
		http.Error(w, message, http.StatusBadRequest)
	}

	if r.Body == nil {
		badRequest("Body should not be empty")
		return
	}

	err := jsonRequestDecode(r.Body, &request)

	if err != nil {
		badRequest(err.Error())
		return
	}

	// set claims to be used by react-app
	response.Claims.User = User{Username: request.Username}
	response.Claims.IssuedAt = time.Now().Unix()
	response.Claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()

	response.Token, _ = generateToken(response.Claims, jwtKey)

	jsonResponseEncode(w, response)
}
