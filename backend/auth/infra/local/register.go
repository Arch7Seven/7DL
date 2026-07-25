package local

import (
	"7DL/auth/domain"
	"fmt"
)

func (auth *Auth) Register(payload *domain.RegisterPayload) error {
	username := payload.Username
	email := payload.Email
	password := payload.Password

	err := ValidateUsername(username,
		auth.Config.UsernameMinLength,
		auth.Config.UsernameMinLength,
	)
	if err != nil {
		return err
	}

	err = ValidatePassword(password,
		auth.Config.PasswordMaxLength,
		auth.Config.PasswordMinLength,
	)
	if err != nil {
		return err
	}

	if payload.Email != nil {
		err = ValidateEmail(*email)
		if err != nil {
			return err
		}
	}

	credentials, err := NewPasswordHash(auth.Config.Hash, password)
	if err != nil {
		return err
	}

	identity, err := auth.Identity.Create(username, email)
	if err != nil {
		return err
	}

	fmt.Println(credentials.PasswordHash)
	fmt.Println(identity.Uuid)

	return nil
}
