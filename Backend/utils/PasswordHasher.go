package utils

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	errors := []string{}
	LCregex := regexp.MustCompile(`^.*[a-z].*$`)
	UCregex := regexp.MustCompile(`^.*[A-Z].*$`)
	Numregex := regexp.MustCompile(`^.*[0-9].*$`)
	Specregex := regexp.MustCompile(`^.*[!@#~$%^&*()].*$`)
	SizeRegex := regexp.MustCompile(`^.{10,}$`)
	AllowedRegex := regexp.MustCompile(`^[A-Za-z\d@$!%*?&]*$`)

	if !LCregex.Match([]byte(password)) {
		errors = append(errors, "\nPassword must contain at least one lowercase letter")
	}
	if !UCregex.Match([]byte(password)) {
		errors = append(errors, "\nPassword must contain at least one uppercase letter")
	}
	if !Numregex.Match([]byte(password)) {
		errors = append(errors, "\nPassword must contain at least one number")
	}
	if !Specregex.Match([]byte(password)) {
		errors = append(errors, "\nPassword must contain at least one special character")
	}
	if !SizeRegex.Match([]byte(password)) {
		errors = append(errors, "\nPassword must be at least 10 characters long")
	}
	if !AllowedRegex.Match([]byte(password)) {
		errors = append(errors, "\nPassword must contain only numbers, letters, and special characters @$!%*?&")
	}

	if len(errors) > 0 {
		return "", fmt.Errorf("Invalid password: %v", strings.Join(errors, ", "))
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Error hashing password: %v", err)
	}

	return string(hashedBytes), nil
}

func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
