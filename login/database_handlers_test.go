package main

import "testing"

func TestCheckCredentials(t *testing.T) {
	_, err := CheckCredentials("testuser@email.com", "testpassword123")

	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
