package app

import (
	"7DL/auth/app"
	"7DL/auth/infra/local"
	"7DL/server"
	"7DL/server/handler"
	"net/http"
)

type App struct {
	HttpServer *http.Server
}

func New() {
	ServerBootstrap()
}

func ServerBootstrap() {
	router := server.NewRouter()
	routes := server.NewRoutes()
	srv := server.New(":8601", router)

	authAdapter := local.New()
	authApp := app.New(authAdapter)
	authHandler := handler.NewAuthHandler(authApp)

	router.Post(routes.Register(), authHandler.Register)
	router.Post(routes.Login(), authHandler.Login)

	srv.ListenAndServe()
}
