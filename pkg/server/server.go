package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server interface {
	Run()
}

type server struct {
	Handler http.Handler
	Port    string
}

func New(handler http.Handler, port string) Server {
	return &server{
		Handler: handler,
		Port:    port,
	}
}

func (server *server) Run() {
	slog.Info(fmt.Sprintf("Server run on %s port", server.Port))

	err := http.ListenAndServe(server.Port, server.Handler)
	if err != nil {
		panic(err.Error())
	}
}
