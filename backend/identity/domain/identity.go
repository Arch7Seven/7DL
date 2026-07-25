package domain

import "github.com/google/uuid"

type Identity struct {
	Uuid     uuid.UUID
	Username string
	Email    *string
}
