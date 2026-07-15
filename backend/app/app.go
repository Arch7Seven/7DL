package app

import (
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

	authHandler := handler.NewAuthHandler()
	router.Post(routes.Register(), authHandler.Register)
	router.Post(routes.Login(), authHandler.Login)

	srv.ListenAndServe()
}
