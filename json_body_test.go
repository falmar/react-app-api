// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bytes"
	"testing"
)

// ------------------ jsonRequestDecode

func TestJSONRequestDecodeMap(t *testing.T) {
	var err error
	var parseInto map[string]string

	body := bytes.NewBuffer([]byte(`{"username": "something", "name": "who knows"}`))

	var expectedValue = map[string]string{
		"username": "something",
		"name":     "who knows",
	}

	// parse
	err = jsonRequestDecode(body, &parseInto)

	if err != nil {
		t.Fatalf("Expected Error to be nil; Got: %v", err)
	}

	// check len
	if len(parseInto) != len(expectedValue) {
		t.Fatalf("Expected map lenth to be %d; Got: %d", len(expectedValue), len(parseInto))
	}

	// check fields
	for i, v := range parseInto {
		if expectedValue[i] != v {
			t.Fatalf("Expected value to be %s; Got: %s", expectedValue[i], v)
		}
	}
}

func TestJSONRequestDecodeStruct(t *testing.T) {
	var err error
	var parseInto struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}

	var expectedValue = struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}{
		"welp", "nope",
	}

	body := bytes.NewBuffer([]byte(`{"username": "welp", "name": "nope"}`))

	err = jsonRequestDecode(body, &parseInto)

	if err != nil {
		t.Fatalf("Expected Error to be nil; Got: %v", err)
	}

	// check property
	if expectedValue.Username != parseInto.Username {
		t.Fatalf("Expected Username to be %s; Got: %s", expectedValue.Username, parseInto.Username)
	}

	// check property
	if expectedValue.Name != parseInto.Name {
		t.Fatalf("Expected Name to be %s; Got: %s", expectedValue.Name, parseInto.Name)
	}
}

func TestJSONRequestDecodeError(t *testing.T) {
	var err error
	var parseInto interface{}

	body := bytes.NewBuffer([]byte(`{"username": "ops error an here, "name": "who knows"}`))

	err = jsonRequestDecode(body, &parseInto)

	if err == nil {
		t.Fatalf("Expected error to be nil; Error: %v", err)
	}
}

// ------------------ End jsonRequestDecode

// ------------------ jsonResponseEncode

func TestJSONResponseEncodeMap(t *testing.T) {
	var err error
	var parseFrom = map[string]string{
		"name":     "iron man. lol!",
		"username": "robocop?",
	}

	var expectedValue = []byte(`{"name":"iron man. lol!","username":"robocop?"}
`)

	body := bytes.NewBuffer([]byte{})

	err = jsonResponseEncode(body, parseFrom)

	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if !bytes.Equal(body.Bytes(), expectedValue) {
		t.Fatalf("Expected body: %s; Got: %s", expectedValue, body.String())
	}
}

func TestJSONResponseEncodeStruct(t *testing.T) {
	var err error
	var parseFrom = struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}{
		"welp", "nope",
	}

	var expectedValue = []byte(`{"username":"welp","name":"nope"}
`)

	body := bytes.NewBuffer([]byte{})

	err = jsonResponseEncode(body, parseFrom)

	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if !bytes.Equal(body.Bytes(), expectedValue) {
		t.Fatalf("Expected body: %s; Got: %s", expectedValue, body.String())
	}
}

// ------------------ End jsonResponseEncode
