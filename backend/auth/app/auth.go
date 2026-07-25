package app

import "7DL/auth/domain"

type AuthApp struct {
	adapter domain.AuthPort
}

func New(adapter domain.AuthPort) *AuthApp {
	return &AuthApp{
		adapter: adapter,
	}
}

func (app *AuthApp) Register(username string, password string, email *string) {
	app.adapter.Register(&domain.RegisterPayload{
		Username: username,
		Password: password,
		Email:    email,
	})
}

func (app *AuthApp) Login(username string, password string) {
	app.adapter.Login(&domain.LoginPayload{
		Username: username,
		Password: password,
	})
}
