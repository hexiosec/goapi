// Code generated by goapi. DO NOT EDIT.
package server

// Error
type Error struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// ListPetsQuery
type ListPetsQuery struct {

	// How many items to return at one time (max 100)
	Limit *int `query:"limit" validate:"omitempty,lte=100"`
}

// Pet: Pet object
type Pet struct {
	ID int `json:"id" validate:"required"`

	// Pet Name
	Name string `json:"name" validate:"required"`

	// List of tags
	Tag *string `json:"tag,omitempty" validate:"omitempty"`
}

// Pets
type Pets []Pet
