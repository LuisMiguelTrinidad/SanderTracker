package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type PasswordRule struct {
	pattern *regexp.Regexp
	message string
}

var passwordRules = []PasswordRule{
	{regexp.MustCompile(`[a-z]`), "lowercase letter"},
	{regexp.MustCompile(`[A-Z]`), "uppercase letter"},
	{regexp.MustCompile(`[0-9]`), "number"},
	{regexp.MustCompile(`[!@#~$%^&*()]`), "special character"},
	{regexp.MustCompile(`^.{10,}$`), "minimum length of 10 characters"},
	{regexp.MustCompile(`^[A-Za-z\d@$!%*?&]*$`), "allowed characters only (A-Z, a-z, 0-9, @$!%*?&)"},
}

func validatePassword(password string) error {
	var failedRules []string

	for _, rule := range passwordRules {
		if !rule.pattern.MatchString(password) {
			failedRules = append(failedRules, rule.message)
		}
	}

	if len(failedRules) > 0 {
		return fmt.Errorf("password must contain: %v", failedRules)
	}
	return nil
}

func HashPassword(password string) (string, error) {
	if err := validatePassword(password); err != nil {
		return "", err
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedBytes), nil
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
