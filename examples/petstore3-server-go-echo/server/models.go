// Code generated by goapi. DO NOT EDIT.
package server

import "time"

// Address
type Address struct {
	City   *string `json:"city,omitempty" validate:"omitempty"`
	State  *string `json:"state,omitempty" validate:"omitempty"`
	Street *string `json:"street,omitempty" validate:"omitempty"`
	Zip    *string `json:"zip,omitempty" validate:"omitempty"`
}

// ApiResponse
type ApiResponse struct {
	Code    *int    `json:"code,omitempty" validate:"omitempty"`
	Message *string `json:"message,omitempty" validate:"omitempty"`
	Type    *string `json:"type,omitempty" validate:"omitempty"`
}

// Category
type Category struct {
	ID   *int    `json:"id,omitempty" validate:"omitempty"`
	Name *string `json:"name,omitempty" validate:"omitempty"`
}

// CreateUsersWithListInputJSONRequest
type CreateUsersWithListInputJSONRequest []User

// Customer
type Customer struct {
	Address  *[]Address `json:"address,omitempty" validate:"omitempty,dive,required"`
	ID       *int       `json:"id,omitempty" validate:"omitempty"`
	Username *string    `json:"username,omitempty" validate:"omitempty"`
}

// FindPetsByStatusJSON200Response
type FindPetsByStatusJSON200Response []Pet

// FindPetsByStatusQuery
type FindPetsByStatusQuery struct {

	// Status values that need to be considered for filter
	Status *string `query:"status" validate:"omitempty,oneof=available pending sold"`
}

// FindPetsByTagsJSON200Response
type FindPetsByTagsJSON200Response []Pet

// FindPetsByTagsQuery
type FindPetsByTagsQuery struct {
	Limit *int `query:"limit" validate:"omitempty,lte=1000"`

	// Tags to filter by
	Tags []string `query:"tags" validate:"omitempty,dive,required"`
}

// GetInventoryJSON200Response
type GetInventoryJSON200Response map[string]interface{}

// LoginUserJSON200Response
type LoginUserJSON200Response string

// LoginUserQuery
type LoginUserQuery struct {

	// The password for login in clear text
	Password *string `query:"password" validate:"omitempty"`

	// The user name for login
	Username *string `query:"username" validate:"omitempty"`
}

// Order
type Order struct {
	Complete *bool      `json:"complete,omitempty" validate:"omitempty"`
	ID       *int       `json:"id,omitempty" validate:"omitempty"`
	PetID    *int       `json:"petId,omitempty" validate:"omitempty"`
	Quantity *int       `json:"quantity,omitempty" validate:"omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty" validate:"omitempty"`

	// Order Status
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=placed approved delivered"`
}

// Pet
type Pet struct {
	Category  *Category `json:"category,omitempty" validate:"omitempty"`
	ID        *int      `json:"id,omitempty" validate:"omitempty,gt=0"`
	Name      string    `json:"name" validate:"required"`
	PhotoURLs []string  `json:"photoUrls" validate:"required,dive,required"`

	// pet status in the store
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=available pending sold"`
	Tags   *[]Tag  `json:"tags,omitempty" validate:"omitempty,dive,required"`
}

// Tag
type Tag struct {
	ID   *int    `json:"id,omitempty" validate:"omitempty"`
	Name *string `json:"name,omitempty" validate:"omitempty"`
}

// UpdatePetWithFormQuery
type UpdatePetWithFormQuery struct {

	// Name of pet that needs to be updated
	Name *string `query:"name" validate:"omitempty"`

	// Status of pet that needs to be updated
	Status *string `query:"status" validate:"omitempty"`
}

// UploadFileQuery
type UploadFileQuery struct {

	// Additional Metadata
	AdditionalMetadata *string `query:"additionalMetadata" validate:"omitempty"`
}

// User
type User struct {
	Email     *string `json:"email,omitempty" validate:"omitempty"`
	FirstName *string `json:"firstName,omitempty" validate:"omitempty"`
	ID        *int    `json:"id,omitempty" validate:"omitempty"`
	LastName  *string `json:"lastName,omitempty" validate:"omitempty"`
	Password  *string `json:"password,omitempty" validate:"omitempty"`
	Phone     *string `json:"phone,omitempty" validate:"omitempty"`

	// User Status
	UserStatus *int    `json:"userStatus,omitempty" validate:"omitempty"`
	Username   *string `json:"username,omitempty" validate:"omitempty"`
}
