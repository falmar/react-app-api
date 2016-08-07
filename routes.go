// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.
//+build !test

package main

import "net/http"

func setRoutes(r *routerWrapper) {
	r.POST("/login", http.HandlerFunc(loginHandler))
}
