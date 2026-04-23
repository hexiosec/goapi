package main

import (
	"example.com/petstore/server"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Set up a basic echo server
	e := echo.New()
	e.Use(middleware.Recover())

	// Create a server RouteHandler from a new Routes struct
	// Routes must implement server.PetsEndpoints
	rh := server.NewPetsRouteHandlers(&Routes{})

	// Register routes at their default endpoints, optionally providing middleware
	rh.RegisterCreatePetsRoute(e)
	rh.RegisterListPetsRoute(e)
	rh.RegisterShowPetByIDRoute(e)

	// Register copy of the create route at a custom endpoint
	rh.RegisterCreatePetsRouteAt("/new", e)

	// Run the server
	if err := e.Start(":3000"); err != nil {
		e.Logger.Error("server stopped", "error", err)
	}
}
