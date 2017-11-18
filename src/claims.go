package main

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims claims of user
type Claims struct {
	jwt.StandardClaims
	User User `json:"user"`
	Role Role `json:"role"`
}
