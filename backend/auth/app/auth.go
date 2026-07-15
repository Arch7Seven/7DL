package app

import "7DL/auth/domain"

// TODO password char max : 100

type AuthApp struct {
	Adapter domain.AuthPort
}

func (app *AuthApp) Register(username string, password string, email *string) {
	app.Adapter.Register(&domain.RegisterPayload{
		Username: username,
		Password: password,
		Email:    email,
	})
}

func (app *AuthApp) Login(username string, password string) {
	app.Adapter.Login(&domain.LoginPayload{
		Username: username,
		Password: password,
	})
}
