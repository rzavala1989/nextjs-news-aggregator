// internal/auth/jwt.go
package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	log.Println("Generating JWT for user:", username)
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error signing JWT:", err)
		return "", err
	}
	return tokenStr, nil
}
