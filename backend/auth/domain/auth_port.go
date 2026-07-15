package domain

type AuthPort interface {
	Register(password *RegisterPayload) error
	Login(payload *LoginPayload) error
}
