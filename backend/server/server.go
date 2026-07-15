package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(addr string, router http.Handler) *http.Server {

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return server
}

func NewRouter() *chi.Mux {
	return chi.NewRouter()
}
