package domain

import "github.com/google/uuid"

type UserCredentials struct {
	UserUUID     uuid.UUID
	Salt         string
	PasswordHash string
}

type LoginPayload struct {
	Username string
	Password string
}

type RegisterPayload struct {
	Username string
	Email    *string
	Password string
}
