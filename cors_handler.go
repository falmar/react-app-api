// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CORS struct {
	router *httprouter.Router
}

func (c CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
		w.Header().Set(
			"Access-Control-Allow-Methods",
			r.Header.Get("Access-Control-Allow-Methods"),
		)
	}

	if r.Method == "OPTIONS" {
		return
	}

	c.router.ServeHTTP(w, r)
}
