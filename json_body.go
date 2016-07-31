// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io"
)

func jsonRequestDecode(r io.ReadCloser, parseInto interface{}) error {
	var decoder *json.Decoder
	var err error

	decoder = json.NewDecoder(r)

	err = decoder.Decode(parseInto)

	if err != nil {
		return err
	}

	return nil
}

func jsonResponseEncode(w io.Writer, parseFrom interface{}) error {
	var encoder *json.Encoder
	var err error

	encoder = json.NewEncoder(w)

	err = encoder.Encode(parseFrom)

	if err != nil {
		return err
	}

	return nil
}
