package handler

import (
	"7DL/auth/app"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	Auth *app.AuthApp
}

func NewAuthHandler(auth *app.AuthApp) *AuthHandler {
	return &AuthHandler{
		Auth: auth,
	}
}

func (handler *AuthHandler) Register(http.ResponseWriter, *http.Request) {
	fmt.Println("a")
	handler.Auth.Register("username", "password", nil)
}

func (handler *AuthHandler) Login(http.ResponseWriter, *http.Request) {
	handler.Auth.Login("username", "password")
}
