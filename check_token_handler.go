// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"
	"time"
)

func checkTokenHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")

	if tokenString == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	token, err := parseToken(tokenString, &MyClaims{}, JWT_KEY)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	claims := token.Claims.(*MyClaims)

	iat := time.Since(time.Unix(claims.IssuedAt, 0))

	if iat.Minutes() >= 90 {
		claims.IssuedAt = time.Now().Unix()

		tokenString, err = generateToken(claims, JWT_KEY)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	jsonResponseEncode(w, LoginResponse{
		Claims: *claims,
		Token:  tokenString,
	})

}
