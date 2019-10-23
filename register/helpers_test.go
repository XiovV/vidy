package main

import "testing"

func TestIsEmailValid(t *testing.T) {
	var passwords = []struct {
		password string
		expected bool
	}{
		{"user.com", false},
		{"user123", false},
		{"user123@password.com", true},
		{"user@password.com", true},
	}

	t.Run("Test if IspasswordValid works properly", func(t *testing.T) {
		for _, password := range passwords {
			if output, _ := IsEmailValid(password.password); output != password.expected {
				t.Errorf("input: %v | expected %v | recieved: %v", password.password, password.expected, output)
			}
		}
	})
}

func TestIsPasswordValid(t *testing.T) {
	var passwords = []struct {
		password string
		expected bool
	}{
		{"abcd123", false},
		{"123", false},
		{"abcdefgh", true},
		{"1234567890", true},
	}

	t.Run("Test if IsPasswordValid works properly", func(t *testing.T) {
		for _, password := range passwords {
			if output, _ := IsPasswordValid(password.password); output != password.expected {
				t.Errorf("input: %v | expected %v | recieved: %v", password.password, password.expected, output)
			}
		}
	})
}

func TestArePasswordsTheSame(t *testing.T) {
	var passwords = []struct {
		password   string
		password_c string
		expected   bool
	}{
		{"abcde123", "abcd123", false},
		{"abcdefgh", "dasdahdasd", false},

		{"abcde123", "abcde123", true},
		{"1234567890", "1234567890", true},
	}

	t.Run("Test if ArePasswordsTheSame works properly", func(t *testing.T) {
		for _, password := range passwords {
			if output, _ := ArePasswordsTheSame(password.password, password.password_c); output != password.expected {
				t.Errorf("input: %v | expected %v | recieved: %v", password.password, password.expected, output)
			}
		}
	})
}
