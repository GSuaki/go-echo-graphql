package main

import (
	"flag"
	"log"
	"os"

	"github.com/gsuaki/go-echo-graphql/api"
)

var envFlag string

func main() {
	flag.StringVar(&envFlag, "env", "server", "environment to run under (gcp, aws, azure, or local)")

	var server *Server
	var err error

	switch envFlag {
	case "test":
		server, err = setupMock()
	default:
		server, err = setupServer()
	}

	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}

type Server struct {
	api *api.Router
}

func NewServer(api *api.Router) *Server {
	return &Server{
		api: api,
	}
}

func (s Server) Start() {
	router := s.api.Router
	router.Logger.Fatal(router.Start(port()))
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	return port
}
