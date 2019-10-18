package main

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// IsEmailValid does very basic email validation
func IsEmailValid(email string) (bool, error) {
	if strings.Contains(email, "@") == true {
		return true, nil
	}
	return false, fmt.Errorf("Email doesn't contain @")
}

// IsPasswordValid does very basic email validation
func IsPasswordValid(password string) (bool, error) {
	if len(password) < 8 {
		return false, fmt.Errorf("Password is less than 8 characters long")
	}
	return true, nil
}
