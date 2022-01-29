//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gsuaki/go-echo-graphql/api"
	"github.com/gsuaki/go-echo-graphql/api/rest"
)

func setupServer() (*Server, error) {
	wire.Build(
		NewServer,
		api.NewRouter,
		rest.NewHealthHandler,
	)
	return &Server{}, nil
}
