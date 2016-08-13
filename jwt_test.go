// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

type badClaim map[interface{}]interface{}

func (bc badClaim) Valid() error {
	return nil
}

func TestGenerateToken(t *testing.T) {
	claims := MyClaims{
		User: &User{
			Username: "dlavieri",
		},
	}

	tokenString, err := generateToken(claims, []byte("super-secret"))

	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if tokenString == "" {
		t.Fatal("Expected token not to be empty")
	}
}

func TestGenerateTokenInvalidClaims(t *testing.T) {
	claims := badClaim{
		1: "not marshable",
	}

	_, err := generateToken(claims, []byte("super-secret"))

	if err == nil {
		t.Fatal("Expected error to be nil")
	}
}

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVzZXJuYW1lIjoiZGxhdmllcmkifX0.dwX1W9X6YSNnRoBIQHWMAMIU-M48UtdSLgXn-TBemP8"

	token, err := parseToken(tokenString, &MyClaims{}, []byte("super-secret"))

	if err != nil {
		t.Fatal(err)
	}

	claims, ok := token.Claims.(*MyClaims)

	if !ok {
		t.Fatalf("Expected claims to be type: %T; Got: %T", MyClaims{}, token.Claims)
	}

	if claims.User.Username != "dlavieri" {
		t.Fatalf("Expected value to be: %s; Got: %v ", "dlavieri", claims.User.Username)
	}
}

func TestParseTokenBadClaim(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7InVzZXJuYW1lIjoiIn19.e0gV7oUokFgLGn7Ht-0mnmh2b3cOC836i4SPaOV7l1c"

	_, err := parseToken(tokenString, &MyClaims{}, []byte("super-secret"))

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}
}

func TestParseTokenBadMethod(t *testing.T) {
	tokenString := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJtYWlsdG86bWlrZUBleGFtcGxlLmNvbSIsIm5iZiI6MTQ3MDYwODc3NywiZXhwIjoxNDcwNjEyMzc3LCJpYXQiOjE0NzA2MDg3NzcsImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.Mhp5qdiadyF6EOlHXUyfMa2aM6QUjU_GiIk7Rfl4CcvhJ5PlgD-aIrhzmW91iNNrM0pAKyOHNsAaV-h3DEiing"

	_, err := parseToken(tokenString, &MyClaims{}, []byte("super-secret"))

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}
}
