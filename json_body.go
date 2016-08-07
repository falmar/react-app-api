// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io"
)

func jsonRequestDecode(r io.Reader, parseInto interface{}) error {
	var err error

	err = json.NewDecoder(r).Decode(parseInto)

	if err != nil {
		return err
	}

	return nil
}

func jsonResponseEncode(w io.Writer, parseFrom interface{}) error {
	var err error

	err = json.NewEncoder(w).Encode(parseFrom)

	if err != nil {
		return err
	}

	return nil
}
