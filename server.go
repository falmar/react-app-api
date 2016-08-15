// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.
//+build !test

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Println("Serving on port", port)

	log.Fatal(
		http.ListenAndServe(port, cors),
	)
}
