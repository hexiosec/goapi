package main

import (
	"log"
	"net/http"

	"github.com/hexiosec/goapi/examples/petstore/server"
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func main() {
	e := echo.New()
	e.Debug = true

	rh := server.NewPetsRouteHandlers(&Controller{})

	rh.RegisterCreatePetsRoute(e, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body := c.Get("body").(*server.Pet)
			log.Printf("Create Pet %#v", body)
			return next(c)
		}
	})
	rh.RegisterListPetsRoute(e)
	rh.RegisterShowPetByIDRoute(e)

	e.Start(":3000")
}

// CreatePets implements server.PetsEndpoints.
func (*Controller) CreatePets(c echo.Context, body *server.Pet) error {
	return c.NoContent(http.StatusCreated)
}

// ListPets implements server.PetsEndpoints.
func (*Controller) ListPets(c echo.Context) (*server.Pets, error) {
	return &server.Pets{
		{
			ID:   1,
			Name: "Gerald",
			Tag:  &[]string{"dog"}[0],
		},
	}, nil
}

// ShowPetByID implements server.PetsEndpoints.
func (*Controller) ShowPetByID(c echo.Context, petID string) (*server.Pet, error) {
	panic("unimplemented")
}
