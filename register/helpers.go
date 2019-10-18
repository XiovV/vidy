package main

import (
	"fmt"
	"strings"
)

// IsEmailValid does very basic email validation
func IsEmailValid(email string) (bool, error) {
	if strings.Contains(email, "@") == true {
		return false, nil
	} else {
		return false, fmt.Errorf("Email doesn't contain @")
	}
}

// IsPasswordValid does very basic email validation
func IsPasswordValid(password string) (bool, error) {
	if len(password) < 8 {
		return false, fmt.Errorf("Password is less than 8 characters long")
	} else {
		return true, nil
	}
}
