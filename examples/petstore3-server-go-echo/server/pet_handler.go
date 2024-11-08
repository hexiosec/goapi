// Code generated by goapi. DO NOT EDIT.
package server

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Interface for Pet route endpoints
type PetEndpoints interface {

	// Add a new pet to the store
	AddPet(c echo.Context, body *Pet) (*Pet, error)

	// Update an existing pet by Id
	UpdatePet(c echo.Context, body *Pet) (*Pet, error)

	// Multiple status values can be provided with comma separated strings
	FindPetsByStatus(c echo.Context, query *FindPetsByStatusQuery) (*FindPetsByStatusJSON200Response, error)

	// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
	FindPetsByTags(c echo.Context, query *FindPetsByTagsQuery) (*FindPetsByTagsJSON200Response, error)
	DeletePet(c echo.Context, xAPIKey string, petID string) error

	// Returns a single pet
	GetPetByID(c echo.Context, petID string) (*Pet, error)
	UpdatePetWithForm(c echo.Context, petID string, query *UpdatePetWithFormQuery) error
	UploadFile(c echo.Context, petID string, query *UploadFileQuery) (*ApiResponse, error)
}

// Wrapper to expose PetEndpoints functions as echo handlers/middleware
type PetRouteHandlers struct {
	validate *validator.Validate
	wrapper  PetEndpoints
}

func NewPetRouteHandlers(wrapper PetEndpoints) *PetRouteHandlers {
	return &PetRouteHandlers{
		validate: validator.New(validator.WithRequiredStructEnabled()),
		wrapper:  wrapper,
	}
}

//------------------------------------------------------------------------------
// # addPet: Add a new pet to the store
//
// POST:/pet
//
// Add a new pet to the store
//
// ## Security
//
// - petstore_auth: write:pets, read:pets
//
// ## Request Body
//
// content:
//     application/json:
//         schema:
//             $ref: '#/components/schemas/Pet'
//     application/x-www-form-urlencoded:
//         schema:
//             $ref: '#/components/schemas/Pet'
//     application/xml:
//         schema:
//             $ref: '#/components/schemas/Pet'
// description: Create a new pet in the store
// required: true
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/Pet'
//         application/xml:
//             schema:
//                 $ref: '#/components/schemas/Pet'
//     description: Successful operation
// "405":
//     description: Invalid input
//------------------------------------------------------------------------------

// Validate requests to POST:/pet
func (r *PetRouteHandlers) AddPetValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Body: Pet
		body := &Pet{}
		if err := (&echo.DefaultBinder{}).BindBody(c, body); err != nil {
			return err
		} else if err := r.validate.Struct(*body); err != nil {
			return err
		}

		c.Set("body", body)
		return next(c)
	}
}

// Handle requests to POST:/pet
func (r *PetRouteHandlers) AddPetHandler(c echo.Context) error {
	body := c.Get("body").(*Pet)

	if res, err := r.wrapper.AddPet(c, body); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for POST:/pet
func (r *PetRouteHandlers) AddPetPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet", trimPrefix[0])
	}
	return "/pet"
}

// Register the handler and middleware for POST:/pet at the default path
func (r *PetRouteHandlers) RegisterAddPetRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterAddPetRouteAt(r.AddPetPath(), e, m...)
}

// Register the handler and middleware for POST:/pet at a custom path
func (r *PetRouteHandlers) RegisterAddPetRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.AddPetValidator}, m...)
	return e.POST(path, r.AddPetHandler, mw...)
}

//------------------------------------------------------------------------------
// # updatePet: Update an existing pet
//
// PUT:/pet
//
// Update an existing pet by Id
//
// ## Security
//
// - petstore_auth: write:pets, read:pets
//
// ## Request Body
//
// content:
//     application/json:
//         schema:
//             $ref: '#/components/schemas/Pet'
//     application/x-www-form-urlencoded:
//         schema:
//             $ref: '#/components/schemas/Pet'
//     application/xml:
//         schema:
//             $ref: '#/components/schemas/Pet'
// description: Update an existent pet in the store
// required: true
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/Pet'
//         application/xml:
//             schema:
//                 $ref: '#/components/schemas/Pet'
//     description: Successful operation
// "400":
//     description: Invalid ID supplied
// "404":
//     description: Pet not found
// "405":
//     description: Validation exception
//------------------------------------------------------------------------------

// Validate requests to PUT:/pet
func (r *PetRouteHandlers) UpdatePetValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Body: Pet
		body := &Pet{}
		if err := (&echo.DefaultBinder{}).BindBody(c, body); err != nil {
			return err
		} else if err := r.validate.Struct(*body); err != nil {
			return err
		}

		c.Set("body", body)
		return next(c)
	}
}

// Handle requests to PUT:/pet
func (r *PetRouteHandlers) UpdatePetHandler(c echo.Context) error {
	body := c.Get("body").(*Pet)

	if res, err := r.wrapper.UpdatePet(c, body); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for PUT:/pet
func (r *PetRouteHandlers) UpdatePetPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet", trimPrefix[0])
	}
	return "/pet"
}

// Register the handler and middleware for PUT:/pet at the default path
func (r *PetRouteHandlers) RegisterUpdatePetRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterUpdatePetRouteAt(r.UpdatePetPath(), e, m...)
}

// Register the handler and middleware for PUT:/pet at a custom path
func (r *PetRouteHandlers) RegisterUpdatePetRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.UpdatePetValidator}, m...)
	return e.PUT(path, r.UpdatePetHandler, mw...)
}

//------------------------------------------------------------------------------
// # findPetsByStatus: Finds Pets by status
//
// GET:/pet/findByStatus
//
// Multiple status values can be provided with comma separated strings
//
// ## Security
//
// - public
// - petstore_auth: write:pets, read:pets
//
// ## Parameters
//
// - description: Status values that need to be considered for filter
//   in: query
//   name: status
//   schema:
//     description: Status values that need to be considered for filter
//     enum:
//         - available
//         - pending
//         - sold
//     type: string
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/FindPetsByStatusJSON200Response'
//         application/xml:
//             schema:
//                 items:
//                     $ref: '#/components/schemas/Pet'
//                 type: array
//     description: successful operation
// "400":
//     description: Invalid status value
//------------------------------------------------------------------------------

// Validate requests to GET:/pet/findByStatus
func (r *PetRouteHandlers) FindPetsByStatusValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.public", []string{})
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Query: FindPetsByStatusQuery
		query := &FindPetsByStatusQuery{}
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, query); err != nil {
			return err
		} else if err := r.validate.Struct(*query); err != nil {
			return err
		}

		c.Set("query", query)
		return next(c)
	}
}

// Handle requests to GET:/pet/findByStatus
func (r *PetRouteHandlers) FindPetsByStatusHandler(c echo.Context) error {
	query := c.Get("query").(*FindPetsByStatusQuery)

	if res, err := r.wrapper.FindPetsByStatus(c, query); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for GET:/pet/findByStatus
func (r *PetRouteHandlers) FindPetsByStatusPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet/findByStatus", trimPrefix[0])
	}
	return "/pet/findByStatus"
}

// Register the handler and middleware for GET:/pet/findByStatus at the default path
func (r *PetRouteHandlers) RegisterFindPetsByStatusRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterFindPetsByStatusRouteAt(r.FindPetsByStatusPath(), e, m...)
}

// Register the handler and middleware for GET:/pet/findByStatus at a custom path
func (r *PetRouteHandlers) RegisterFindPetsByStatusRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.FindPetsByStatusValidator}, m...)
	return e.GET(path, r.FindPetsByStatusHandler, mw...)
}

//------------------------------------------------------------------------------
// # findPetsByTags: Finds Pets by tags
//
// GET:/pet/findByTags
//
// Multiple tags can be provided with comma separated strings. Use tag1, tag2,
// tag3 for testing.
//
// ## Security
//
// - petstore_auth: write:pets, read:pets
//
// ## Parameters
//
// - description: Tags to filter by
//   in: query
//   name: tags
//   schema:
//     description: Tags to filter by
//     items:
//         type: string
//     type: array
// - in: query
//   name: limit
//   schema:
//     maximum: 1000
//     type: integer
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/FindPetsByTagsJSON200Response'
//         application/xml:
//             schema:
//                 items:
//                     $ref: '#/components/schemas/Pet'
//                 type: array
//     description: successful operation
// "400":
//     description: Invalid tag value
//------------------------------------------------------------------------------

// Validate requests to GET:/pet/findByTags
func (r *PetRouteHandlers) FindPetsByTagsValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Query: FindPetsByTagsQuery
		query := &FindPetsByTagsQuery{}
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, query); err != nil {
			return err
		} else if err := r.validate.Struct(*query); err != nil {
			return err
		}

		c.Set("query", query)
		return next(c)
	}
}

// Handle requests to GET:/pet/findByTags
func (r *PetRouteHandlers) FindPetsByTagsHandler(c echo.Context) error {
	query := c.Get("query").(*FindPetsByTagsQuery)

	if res, err := r.wrapper.FindPetsByTags(c, query); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for GET:/pet/findByTags
func (r *PetRouteHandlers) FindPetsByTagsPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet/findByTags", trimPrefix[0])
	}
	return "/pet/findByTags"
}

// Register the handler and middleware for GET:/pet/findByTags at the default path
func (r *PetRouteHandlers) RegisterFindPetsByTagsRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterFindPetsByTagsRouteAt(r.FindPetsByTagsPath(), e, m...)
}

// Register the handler and middleware for GET:/pet/findByTags at a custom path
func (r *PetRouteHandlers) RegisterFindPetsByTagsRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.FindPetsByTagsValidator}, m...)
	return e.GET(path, r.FindPetsByTagsHandler, mw...)
}

//------------------------------------------------------------------------------
// # deletePet: Deletes a pet
//
// DELETE:/pet/:petId
//
// ## Security
//
// - petstore_auth: write:pets, read:pets
//
// ## Parameters
//
// - in: header
//   name: X-Api-Key
//   required: true
//   schema:
//     type: string
// - description: Pet id to delete
//   in: path
//   name: petId
//   required: true
//   schema:
//     format: int64
//     type: integer
//
// ## Responses
//
// "400":
//     description: Invalid pet value
//------------------------------------------------------------------------------

// Validate requests to DELETE:/pet/:petId
func (r *PetRouteHandlers) DeletePetValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Header Parameter: X-Api-Key
		xAPIKey := c.Request().Header.Get("X-Api-Key")
		if err := r.validate.Var(xAPIKey, "required"); err != nil {
			return err
		}

		c.Set("param.x_api_key", xAPIKey)

		// Path Parameter: petId
		petID := c.Param("petId")
		if err := r.validate.Var(petID, "required"); err != nil {
			return err
		}

		c.Set("param.pet_id", petID)

		return next(c)
	}
}

// Handle requests to DELETE:/pet/:petId
func (r *PetRouteHandlers) DeletePetHandler(c echo.Context) error {
	xAPIKey := c.Get("param.x_api_key").(string)
	petID := c.Get("param.pet_id").(string)

	if err := r.wrapper.DeletePet(c, xAPIKey, petID); err == nil {
		if !c.Response().Committed {
			return c.NoContent(http.StatusNoContent)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for DELETE:/pet/:petId
func (r *PetRouteHandlers) DeletePetPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet/:petId", trimPrefix[0])
	}
	return "/pet/:petId"
}

// Register the handler and middleware for DELETE:/pet/:petId at the default path
func (r *PetRouteHandlers) RegisterDeletePetRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterDeletePetRouteAt(r.DeletePetPath(), e, m...)
}

// Register the handler and middleware for DELETE:/pet/:petId at a custom path
func (r *PetRouteHandlers) RegisterDeletePetRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.DeletePetValidator}, m...)
	return e.DELETE(path, r.DeletePetHandler, mw...)
}

//------------------------------------------------------------------------------
// # getPetById: Find pet by ID
//
// GET:/pet/:petId
//
// Returns a single pet
//
// ## Security
//
// - api_key
// - petstore_auth: write:pets, read:pets
//
// ## Parameters
//
// - description: ID of pet to return
//   in: path
//   name: petId
//   required: true
//   schema:
//     format: int64
//     type: integer
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/Pet'
//         application/xml:
//             schema:
//                 $ref: '#/components/schemas/Pet'
//     description: successful operation
// "400":
//     description: Invalid ID supplied
// "404":
//     description: Pet not found
//------------------------------------------------------------------------------

// Validate requests to GET:/pet/:petId
func (r *PetRouteHandlers) GetPetByIDValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.api_key", []string{})

		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Path Parameter: petId
		petID := c.Param("petId")
		if err := r.validate.Var(petID, "required"); err != nil {
			return err
		}

		c.Set("param.pet_id", petID)

		return next(c)
	}
}

// Handle requests to GET:/pet/:petId
func (r *PetRouteHandlers) GetPetByIDHandler(c echo.Context) error {
	petID := c.Get("param.pet_id").(string)

	if res, err := r.wrapper.GetPetByID(c, petID); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for GET:/pet/:petId
func (r *PetRouteHandlers) GetPetByIDPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet/:petId", trimPrefix[0])
	}
	return "/pet/:petId"
}

// Register the handler and middleware for GET:/pet/:petId at the default path
func (r *PetRouteHandlers) RegisterGetPetByIDRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterGetPetByIDRouteAt(r.GetPetByIDPath(), e, m...)
}

// Register the handler and middleware for GET:/pet/:petId at a custom path
func (r *PetRouteHandlers) RegisterGetPetByIDRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.GetPetByIDValidator}, m...)
	return e.GET(path, r.GetPetByIDHandler, mw...)
}

//------------------------------------------------------------------------------
// # updatePetWithForm: Updates a pet in the store with form data
//
// POST:/pet/:petId
//
// ## Security
//
// - petstore_auth: write:pets, read:pets
//
// ## Parameters
//
// - description: ID of pet that needs to be updated
//   in: path
//   name: petId
//   required: true
//   schema:
//     format: int64
//     type: integer
// - description: Name of pet that needs to be updated
//   in: query
//   name: name
//   schema:
//     description: Name of pet that needs to be updated
//     type: string
// - description: Status of pet that needs to be updated
//   in: query
//   name: status
//   schema:
//     description: Status of pet that needs to be updated
//     type: string
//
// ## Responses
//
// "405":
//     description: Invalid input
//------------------------------------------------------------------------------

// Validate requests to POST:/pet/:petId
func (r *PetRouteHandlers) UpdatePetWithFormValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Path Parameter: petId
		petID := c.Param("petId")
		if err := r.validate.Var(petID, "required"); err != nil {
			return err
		}

		c.Set("param.pet_id", petID)

		// Query: UpdatePetWithFormQuery
		query := &UpdatePetWithFormQuery{}
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, query); err != nil {
			return err
		} else if err := r.validate.Struct(*query); err != nil {
			return err
		}

		c.Set("query", query)
		return next(c)
	}
}

// Handle requests to POST:/pet/:petId
func (r *PetRouteHandlers) UpdatePetWithFormHandler(c echo.Context) error {
	petID := c.Get("param.pet_id").(string)
	query := c.Get("query").(*UpdatePetWithFormQuery)

	if err := r.wrapper.UpdatePetWithForm(c, petID, query); err == nil {
		if !c.Response().Committed {
			return c.NoContent(http.StatusNoContent)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for POST:/pet/:petId
func (r *PetRouteHandlers) UpdatePetWithFormPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet/:petId", trimPrefix[0])
	}
	return "/pet/:petId"
}

// Register the handler and middleware for POST:/pet/:petId at the default path
func (r *PetRouteHandlers) RegisterUpdatePetWithFormRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterUpdatePetWithFormRouteAt(r.UpdatePetWithFormPath(), e, m...)
}

// Register the handler and middleware for POST:/pet/:petId at a custom path
func (r *PetRouteHandlers) RegisterUpdatePetWithFormRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.UpdatePetWithFormValidator}, m...)
	return e.POST(path, r.UpdatePetWithFormHandler, mw...)
}

//------------------------------------------------------------------------------
// # uploadFile: uploads an image
//
// POST:/pet/:petId/uploadImage
//
// ## Security
//
// - petstore_auth: write:pets, read:pets
//
// ## Parameters
//
// - description: ID of pet to update
//   in: path
//   name: petId
//   required: true
//   schema:
//     format: int64
//     type: integer
// - description: Additional Metadata
//   in: query
//   name: additionalMetadata
//   schema:
//     description: Additional Metadata
//     type: string
//
// ## Request Body
//
// content:
//     application/octet-stream:
//         schema:
//             format: binary
//             type: string
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/ApiResponse'
//     description: successful operation
//------------------------------------------------------------------------------

// Validate requests to POST:/pet/:petId/uploadImage
func (r *PetRouteHandlers) UploadFileValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("security.petstore_auth", []string{"write:pets", "read:pets"})

		// Path Parameter: petId
		petID := c.Param("petId")
		if err := r.validate.Var(petID, "required"); err != nil {
			return err
		}

		c.Set("param.pet_id", petID)

		// Query: UploadFileQuery
		query := &UploadFileQuery{}
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, query); err != nil {
			return err
		} else if err := r.validate.Struct(*query); err != nil {
			return err
		}

		c.Set("query", query)
		return next(c)
	}
}

// Handle requests to POST:/pet/:petId/uploadImage
func (r *PetRouteHandlers) UploadFileHandler(c echo.Context) error {
	petID := c.Get("param.pet_id").(string)
	query := c.Get("query").(*UploadFileQuery)

	if res, err := r.wrapper.UploadFile(c, petID, query); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for POST:/pet/:petId/uploadImage
func (r *PetRouteHandlers) UploadFilePath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/pet/:petId/uploadImage", trimPrefix[0])
	}
	return "/pet/:petId/uploadImage"
}

// Register the handler and middleware for POST:/pet/:petId/uploadImage at the default path
func (r *PetRouteHandlers) RegisterUploadFileRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterUploadFileRouteAt(r.UploadFilePath(), e, m...)
}

// Register the handler and middleware for POST:/pet/:petId/uploadImage at a custom path
func (r *PetRouteHandlers) RegisterUploadFileRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.UploadFileValidator}, m...)
	return e.POST(path, r.UploadFileHandler, mw...)
}
