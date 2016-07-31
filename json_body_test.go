// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// -------------- RequestBodyMock
type RequestBodyMock struct {
	body []byte
}

func (frb RequestBodyMock) Read(p []byte) (int, error) {
	var n int

	n = copy(p, frb.body)

	return n, nil
}

func (frb RequestBodyMock) Close() error {
	return nil
}

// ------------------ End Mock

// -------------- RequestBodyMock
type ResponseBodyMock struct {
	body []byte
}

func (frb *ResponseBodyMock) Write(p []byte) (int, error) {
	frb.body = p

	return len(p), nil
}

// ------------------ End Mock

// ------------------ jsonRequestDecode

func TestJSONRequestDecodeMap(t *testing.T) {
	var err error
	var body RequestBodyMock
	var parseInto map[string]string

	body.body = []byte(`{"username": "something", "name": "who knows"}`)

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
	var body RequestBodyMock
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

	body.body = []byte(`{"username": "welp", "name": "nope"}`)

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
	var body RequestBodyMock
	var parseInto interface{}

	body.body = []byte(`{"username": "ops error an here, "name": "who knows"}`)

	err = jsonRequestDecode(body, &parseInto)

	if err == nil {
		t.Fatalf("Expected error to be nil; Error: %v", err)
	}
}

// ------------------ End jsonRequestDecode

// ------------------ jsonResponseEncode

func TestJSONResponseEncodeMap(t *testing.T) {
	var err error
	var body = &ResponseBodyMock{}
	var parseFrom = map[string]string{
		"username": "robocop?",
		"name":     "iron man. lol!",
	}

	err = jsonResponseEncode(body, parseFrom)

	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if len(body.body) == 0 {
		t.Fatal("Body should not be empty")
	}
}

func TestJSONResponseEncodeStruct(t *testing.T) {
	var err error
	var body ResponseBodyMock
	var parseFrom = struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}{
		"welp", "nope",
	}

	err = jsonResponseEncode(&body, parseFrom)

	if err != nil {
		t.Fatalf("Unexpected Error: %v", err)
	}

	if len(body.body) == 0 {
		t.Fatal("Body should not be empty")
	}
}

// ------------------ End jsonResponseEncode
