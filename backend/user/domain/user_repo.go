package user

import "github.com/google/uuid"

type UserRepository interface {
	GetUserByUUID(uuid.UUID) (*User, error)
	GetUserByUsername(string) (*User, error)

	UserExistFromUUID(uuid.UUID) (bool, error)
	UserExistFromUsername(string) (bool, error)

	Save(*User) error
}
