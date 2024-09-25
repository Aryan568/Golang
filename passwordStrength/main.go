package main

import (
	"fmt"
	"regexp"
)

func passwordStrength(password string) string {

	if len(password) <= 8 {
		return "Weak"
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	if hasLower && hasUpper && hasNumber && hasSpecial {
		return "Very Strong"
	}
	if (hasLower && hasUpper && hasNumber) || (hasUpper && hasNumber && hasSpecial) {
		return "Strong"
	}
	return "Moderate"
}

func main() {
	password := "Aryan9978"
	strength := passwordStrength(password)
	fmt.Printf("Password: %s, Strength: %s\n", password, strength)
}
