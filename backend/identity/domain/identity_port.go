package domain

import "github.com/google/uuid"

type IdentityPort interface {
	Create(username string, email *string) (*Identity, error)

	FindByUsername(username string) (*Identity, error)
	FindByEmail(email string) (*Identity, error)
	FindByUuid(uuid uuid.UUID) (*Identity, error)
}
