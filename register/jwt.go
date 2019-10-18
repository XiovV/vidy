package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var hmacSecret = []byte("secretkey123")

func GenerateToken(user string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

// ValidateToken checks if the right user is using the right token
func ValidateToken(tokenString, user string) bool {
	var isValid bool

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user"] == user {
			isValid = true
			return true
		}

		fmt.Println(claims["user"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return isValid
}
