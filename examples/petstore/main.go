package main

import (
	"github.com/hexiosec/goapi/examples/petstore/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Set up a basic echo server
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Create a server RouteHandler from a new Routes struct
	// Routes must implement server.PetsEndpoints
	rh := server.NewPetsRouteHandlers(&Routes{})

	// Register routes at their default endpoints, optionally providing middleware
	rh.RegisterCreatePetsRoute(e, DebugCreatePetsMiddleware)
	rh.RegisterListPetsRoute(e)
	rh.RegisterShowPetByIDRoute(e, CheckShowPetIDMiddleware)

	// Register copy of the create route at a custom endpoint
	rh.RegisterCreatePetsRouteAt(e, "/new", DebugCreatePetsMiddleware)

	// Run the server
	e.Start(":3000")
}
