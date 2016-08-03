// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.
//+build !test

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	serverHandler()
}

func serverHandler() {
	routerWrapper := newWrapper(
		httprouter.New(),
	)

	// routes.go
	setRoutes(routerWrapper)

	// cros_handler.go
	cors := CORS{routerWrapper}

	log.Fatal(
		http.ListenAndServe(":8080", cors),
	)
}
