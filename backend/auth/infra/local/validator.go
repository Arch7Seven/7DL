package local

import (
	"7DL/auth/domain"
	"fmt"
	"net/mail"
	"unicode/utf8"
)

func ValidateUsername(username string, maxLen int, minLen int) error {
	usernameLen := utf8.RuneCountInString(username)

	if usernameLen <= maxLen && usernameLen >= minLen {
		return nil
	}

	return &domain.InvalidUsername{
		Username: username,
		Msg:      fmt.Sprintf("Unexpected email length, max expected: min %d, max %d, received: %d", minLen, maxLen, usernameLen),
		Cause:    "invalid username length",
	}
}

func ValidateEmail(email string) error {
	addr, err := mail.ParseAddress(email)
	if err == nil || addr.Address == email {
		return nil
	}

	return &domain.InvalidEmail{
		Email: email,
		Msg:   fmt.Sprintf("Unexpected email: %s", email),
	}
}

func ValidatePassword(password string, maxLen int, minLen int) error {
	passwordLen := utf8.RuneCountInString(password)

	if passwordLen <= maxLen && passwordLen >= minLen {
		return nil
	}

	return &domain.InvalidPassword{
		Msg:   fmt.Sprintf("Unexpected email length, expected: min %d, max: %d", minLen, maxLen),
		Cause: "invalid password length",
	}
}
