// internal/models/claims.go
package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	jwt.StandardClaims
}
