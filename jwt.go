// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

// MyClaims custom claim to use in this app
type MyClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}

// Valid implement interface jwt.Claims
func (mc MyClaims) Valid() error {
	if mc.User.Username == "" {
		return errors.New("User's name can not be empty")
	}

	return nil
}

func generateToken(claims jwt.Claims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func parseToken(tokenString string, claims jwt.Claims, secretKey []byte) (*jwt.Token, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	}

	return jwt.ParseWithClaims(tokenString, claims, keyFunc)
}
