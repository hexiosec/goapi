package main

//go:generate go run ../../cmd/generate -i ../petstore.yml -o ./server -t server/go-echo --templates-path ../../templates
//go:generate go run golang.org/x/tools/cmd/goimports@latest -w ./server

import (
	"log"
	"net/http"

	"example.com/petstore/server"
	"github.com/labstack/echo/v4"
)

type Routes struct{}

// DebugCreatePetsMiddleware prints the submitted body before the handler runs
func DebugCreatePetsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		body := c.Get("body").(*server.Pet)
		log.Printf("Create Pet %#v", body)
		return next(c)
	}
}

// CheckShowPetIDMiddleware ensures the ID is a known pet
func CheckShowPetIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		petID := c.Get("param.petId").(string)
		if petID != "1" {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return next(c)
	}
}

// CreatePets implements server.PetsEndpoints.
func (*Routes) CreatePets(c echo.Context, body *server.Pet) error {
	return nil
}

// ListPets implements server.PetsEndpoints.
func (*Routes) ListPets(c echo.Context, query *server.ListPetsQuery) (*server.Pets, error) {
	return &server.Pets{
		{
			ID:   1,
			Name: "Gerald",
			Tag:  &[]string{"dog"}[0],
		},
		{
			ID:   2,
			Name: "Harold",
			Tag:  &[]string{"cat"}[0],
		},
	}, nil
}

// ShowPetByID implements server.PetsEndpoints.
func (*Routes) ShowPetByID(c echo.Context, petID string) (*server.Pet, error) {
	return &server.Pet{
		ID:   1,
		Name: "Gerald",
		Tag:  &[]string{"dog"}[0],
	}, nil
}
