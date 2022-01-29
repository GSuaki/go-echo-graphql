package api

import (
	"github.com/gsuaki/go-echo-graphql/api/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	Router *echo.Echo
}

func NewRouter(
	health *rest.HealthHandler,
) *Router {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", health.Ping)

	return &Router{
		Router: e,
	}
}
