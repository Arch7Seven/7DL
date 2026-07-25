package domain

import (
	"7DL/errors"
	"fmt"
)

// InvalidUsername
type InvalidUsername struct {
	Username string
	Msg      string
	Cause    string
}

func (e InvalidUsername) Error() string {
	return e.Msg
}

func (e InvalidUsername) PublicData() errors.PublicErrorData {
	return errors.PublicErrorData{
		Code:      "invalid_username",
		PublicMsg: fmt.Sprintf("Invalid username: %s", e.Cause),
	}
}

// PCause
type InvalidEmail struct {
	Email string
	Msg   string
	Cause string
}

func (e InvalidEmail) Error() string {
	return e.Msg
}

func (e InvalidEmail) PublicData() errors.PublicErrorData {
	return errors.PublicErrorData{
		Code:      "invalid_email",
		PublicMsg: fmt.Sprintf("Invalid email: %s", e.Cause),
	}
}

// InvalidPassword
type InvalidPassword struct {
	Msg   string
	Cause string
}

func (e InvalidPassword) Error() string {
	return e.Msg
}

func (e InvalidPassword) PublicData() errors.PublicErrorData {
	return errors.PublicErrorData{
		Code:      "invalid_password",
		PublicMsg: fmt.Sprintf("Invalid cause: %s", e.Cause),
	}
}
