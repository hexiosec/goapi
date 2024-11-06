// Code generated by goapi. DO NOT EDIT.
package server

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Interface for User route endpoints
type UserEndpoints interface {

	// This can only be done by the logged in user.
	CreateUser(c echo.Context, body *User) error

	// Creates list of users with given input array
	CreateUsersWithListInput(c echo.Context, body *CreateUsersWithListInputJSONRequest) (*User, error)
	LoginUser(c echo.Context, query *LoginUserQuery) (*LoginUserJSON200Response, error)
	LogoutUser(c echo.Context) error

	// This can only be done by the logged in user.
	DeleteUser(c echo.Context, username string) error
	GetUserByName(c echo.Context, username string) (*User, error)

	// This can only be done by the logged in user.
	UpdateUser(c echo.Context, username string, body *User) error
}

// Wrapper to expose UserEndpoints functions as echo handlers/middleware
type UserRouteHandlers struct {
	validate *validator.Validate
	wrapper  UserEndpoints
}

func NewUserRouteHandlers(wrapper UserEndpoints) *UserRouteHandlers {
	return &UserRouteHandlers{
		validate: validator.New(validator.WithRequiredStructEnabled()),
		wrapper:  wrapper,
	}
}

//------------------------------------------------------------------------------
// # createUser: Create user
//
// POST:/user
//
// This can only be done by the logged in user.
//
// ## Request Body
//
// content:
//     application/json:
//         schema:
//             $ref: '#/components/schemas/User'
//     application/x-www-form-urlencoded:
//         schema:
//             $ref: '#/components/schemas/User'
//     application/xml:
//         schema:
//             $ref: '#/components/schemas/User'
// description: Created user object
//
// ## Responses
//
// default:
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/User'
//         application/xml:
//             schema:
//                 $ref: '#/components/schemas/User'
//     description: successful operation
//------------------------------------------------------------------------------

// Validate requests to POST:/user
func (r *UserRouteHandlers) CreateUserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Body: User
		body := &User{}
		if err := (&echo.DefaultBinder{}).BindBody(c, body); err != nil {
			return err
		} else if err := r.validate.Struct(*body); err != nil {
			return err
		}

		c.Set("body", body)
		return next(c)
	}
}

// Handle requests to POST:/user
func (r *UserRouteHandlers) CreateUserHandler(c echo.Context) error {
	body := c.Get("body").(*User)

	if err := r.wrapper.CreateUser(c, body); err == nil {
		if !c.Response().Committed {
			return c.NoContent(http.StatusNoContent)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for POST:/user
func (r *UserRouteHandlers) CreateUserPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user", trimPrefix[0])
	}
	return "/user"
}

// Register the handler and middleware for POST:/user at the default path
func (r *UserRouteHandlers) RegisterCreateUserRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterCreateUserRouteAt(r.CreateUserPath(), e, m...)
}

// Register the handler and middleware for POST:/user at a custom path
func (r *UserRouteHandlers) RegisterCreateUserRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.CreateUserValidator}, m...)
	return e.POST(path, r.CreateUserHandler, mw...)
}

//------------------------------------------------------------------------------
// # createUsersWithListInput: Creates list of users with given input array
//
// POST:/user/createWithList
//
// Creates list of users with given input array
//
// ## Request Body
//
// content:
//     application/json:
//         schema:
//             $ref: '#/components/schemas/CreateUsersWithListInputJSONRequest'
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/User'
//         application/xml:
//             schema:
//                 $ref: '#/components/schemas/User'
//     description: Successful operation
// default:
//     description: successful operation
//------------------------------------------------------------------------------

// Validate requests to POST:/user/createWithList
func (r *UserRouteHandlers) CreateUsersWithListInputValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Body: CreateUsersWithListInputJSONRequest
		body := &CreateUsersWithListInputJSONRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, body); err != nil {
			return err
		} else if err := r.validate.Struct(*body); err != nil {
			return err
		}

		c.Set("body", body)
		return next(c)
	}
}

// Handle requests to POST:/user/createWithList
func (r *UserRouteHandlers) CreateUsersWithListInputHandler(c echo.Context) error {
	body := c.Get("body").(*CreateUsersWithListInputJSONRequest)

	if res, err := r.wrapper.CreateUsersWithListInput(c, body); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for POST:/user/createWithList
func (r *UserRouteHandlers) CreateUsersWithListInputPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user/createWithList", trimPrefix[0])
	}
	return "/user/createWithList"
}

// Register the handler and middleware for POST:/user/createWithList at the default path
func (r *UserRouteHandlers) RegisterCreateUsersWithListInputRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterCreateUsersWithListInputRouteAt(r.CreateUsersWithListInputPath(), e, m...)
}

// Register the handler and middleware for POST:/user/createWithList at a custom path
func (r *UserRouteHandlers) RegisterCreateUsersWithListInputRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.CreateUsersWithListInputValidator}, m...)
	return e.POST(path, r.CreateUsersWithListInputHandler, mw...)
}

//------------------------------------------------------------------------------
// # loginUser: Logs user into the system
//
// GET:/user/login
//
// ## Parameters
//
// - description: The user name for login
//   in: query
//   name: username
//   schema:
//     description: The user name for login
//     type: string
// - description: The password for login in clear text
//   in: query
//   name: password
//   schema:
//     description: The password for login in clear text
//     type: string
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/LoginUserJSON200Response'
//         application/xml:
//             schema:
//                 type: string
//     description: successful operation
//     headers:
//         X-Expires-After:
//             description: date in UTC when token expires
//             schema:
//                 format: date-time
//                 type: string
//         X-Rate-Limit:
//             description: calls per hour allowed by the user
//             schema:
//                 format: int32
//                 type: integer
// "400":
//     description: Invalid username/password supplied
//------------------------------------------------------------------------------

// Validate requests to GET:/user/login
func (r *UserRouteHandlers) LoginUserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Query: LoginUserQuery
		query := &LoginUserQuery{}
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, query); err != nil {
			return err
		} else if err := r.validate.Struct(*query); err != nil {
			return err
		}

		c.Set("query", query)
		return next(c)
	}
}

// Handle requests to GET:/user/login
func (r *UserRouteHandlers) LoginUserHandler(c echo.Context) error {
	query := c.Get("query").(*LoginUserQuery)

	if res, err := r.wrapper.LoginUser(c, query); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for GET:/user/login
func (r *UserRouteHandlers) LoginUserPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user/login", trimPrefix[0])
	}
	return "/user/login"
}

// Register the handler and middleware for GET:/user/login at the default path
func (r *UserRouteHandlers) RegisterLoginUserRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterLoginUserRouteAt(r.LoginUserPath(), e, m...)
}

// Register the handler and middleware for GET:/user/login at a custom path
func (r *UserRouteHandlers) RegisterLoginUserRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.LoginUserValidator}, m...)
	return e.GET(path, r.LoginUserHandler, mw...)
}

//------------------------------------------------------------------------------
// # logoutUser: Logs out current logged in user session
//
// GET:/user/logout
//
// ## Responses
//
// default:
//     description: successful operation
//------------------------------------------------------------------------------

// Validate requests to GET:/user/logout
func (r *UserRouteHandlers) LogoutUserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// Handle requests to GET:/user/logout
func (r *UserRouteHandlers) LogoutUserHandler(c echo.Context) error {

	if err := r.wrapper.LogoutUser(c); err == nil {
		if !c.Response().Committed {
			return c.NoContent(http.StatusNoContent)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for GET:/user/logout
func (r *UserRouteHandlers) LogoutUserPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user/logout", trimPrefix[0])
	}
	return "/user/logout"
}

// Register the handler and middleware for GET:/user/logout at the default path
func (r *UserRouteHandlers) RegisterLogoutUserRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterLogoutUserRouteAt(r.LogoutUserPath(), e, m...)
}

// Register the handler and middleware for GET:/user/logout at a custom path
func (r *UserRouteHandlers) RegisterLogoutUserRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.LogoutUserValidator}, m...)
	return e.GET(path, r.LogoutUserHandler, mw...)
}

//------------------------------------------------------------------------------
// # deleteUser: Delete user
//
// DELETE:/user/:username
//
// This can only be done by the logged in user.
//
// ## Parameters
//
// - description: The name that needs to be deleted
//   in: path
//   name: username
//   required: true
//   schema:
//     type: string
//
// ## Responses
//
// "400":
//     description: Invalid username supplied
// "404":
//     description: User not found
//------------------------------------------------------------------------------

// Validate requests to DELETE:/user/:username
func (r *UserRouteHandlers) DeleteUserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Path Parameter: username
		username := c.Param("username")
		if err := r.validate.Var(username, "required"); err != nil {
			return err
		}

		c.Set("param.username", username)

		return next(c)
	}
}

// Handle requests to DELETE:/user/:username
func (r *UserRouteHandlers) DeleteUserHandler(c echo.Context) error {
	username := c.Get("param.username").(string)

	if err := r.wrapper.DeleteUser(c, username); err == nil {
		if !c.Response().Committed {
			return c.NoContent(http.StatusNoContent)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for DELETE:/user/:username
func (r *UserRouteHandlers) DeleteUserPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user/:username", trimPrefix[0])
	}
	return "/user/:username"
}

// Register the handler and middleware for DELETE:/user/:username at the default path
func (r *UserRouteHandlers) RegisterDeleteUserRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterDeleteUserRouteAt(r.DeleteUserPath(), e, m...)
}

// Register the handler and middleware for DELETE:/user/:username at a custom path
func (r *UserRouteHandlers) RegisterDeleteUserRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.DeleteUserValidator}, m...)
	return e.DELETE(path, r.DeleteUserHandler, mw...)
}

//------------------------------------------------------------------------------
// # getUserByName: Get user by user name
//
// GET:/user/:username
//
// ## Parameters
//
// - description: 'The name that needs to be fetched. Use user1 for testing. '
//   in: path
//   name: username
//   required: true
//   schema:
//     type: string
//
// ## Responses
//
// "200":
//     content:
//         application/json:
//             schema:
//                 $ref: '#/components/schemas/User'
//         application/xml:
//             schema:
//                 $ref: '#/components/schemas/User'
//     description: successful operation
// "400":
//     description: Invalid username supplied
// "404":
//     description: User not found
//------------------------------------------------------------------------------

// Validate requests to GET:/user/:username
func (r *UserRouteHandlers) GetUserByNameValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Path Parameter: username
		username := c.Param("username")
		if err := r.validate.Var(username, "required"); err != nil {
			return err
		}

		c.Set("param.username", username)

		return next(c)
	}
}

// Handle requests to GET:/user/:username
func (r *UserRouteHandlers) GetUserByNameHandler(c echo.Context) error {
	username := c.Get("param.username").(string)

	if res, err := r.wrapper.GetUserByName(c, username); err == nil {
		if !c.Response().Committed {
			return c.JSON(200, res)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for GET:/user/:username
func (r *UserRouteHandlers) GetUserByNamePath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user/:username", trimPrefix[0])
	}
	return "/user/:username"
}

// Register the handler and middleware for GET:/user/:username at the default path
func (r *UserRouteHandlers) RegisterGetUserByNameRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterGetUserByNameRouteAt(r.GetUserByNamePath(), e, m...)
}

// Register the handler and middleware for GET:/user/:username at a custom path
func (r *UserRouteHandlers) RegisterGetUserByNameRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.GetUserByNameValidator}, m...)
	return e.GET(path, r.GetUserByNameHandler, mw...)
}

//------------------------------------------------------------------------------
// # updateUser: Update user
//
// PUT:/user/:username
//
// This can only be done by the logged in user.
//
// ## Parameters
//
// - description: name that needs to be updated
//   in: path
//   name: username
//   required: true
//   schema:
//     type: string
//
// ## Request Body
//
// content:
//     application/json:
//         schema:
//             $ref: '#/components/schemas/User'
//     application/x-www-form-urlencoded:
//         schema:
//             $ref: '#/components/schemas/User'
//     application/xml:
//         schema:
//             $ref: '#/components/schemas/User'
// description: Update an existent user in the store
//
// ## Responses
//
// default:
//     description: successful operation
//------------------------------------------------------------------------------

// Validate requests to PUT:/user/:username
func (r *UserRouteHandlers) UpdateUserValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Path Parameter: username
		username := c.Param("username")
		if err := r.validate.Var(username, "required"); err != nil {
			return err
		}

		c.Set("param.username", username)

		// Body: User
		body := &User{}
		if err := (&echo.DefaultBinder{}).BindBody(c, body); err != nil {
			return err
		} else if err := r.validate.Struct(*body); err != nil {
			return err
		}

		c.Set("body", body)
		return next(c)
	}
}

// Handle requests to PUT:/user/:username
func (r *UserRouteHandlers) UpdateUserHandler(c echo.Context) error {
	username := c.Get("param.username").(string)
	body := c.Get("body").(*User)

	if err := r.wrapper.UpdateUser(c, username, body); err == nil {
		if !c.Response().Committed {
			return c.NoContent(http.StatusNoContent)
		} else {
			return nil
		}
	} else {
		return err
	}
}

// Get path for PUT:/user/:username
func (r *UserRouteHandlers) UpdateUserPath(trimPrefix ...string) string {
	if len(trimPrefix) > 0 {
		return strings.TrimPrefix("/user/:username", trimPrefix[0])
	}
	return "/user/:username"
}

// Register the handler and middleware for PUT:/user/:username at the default path
func (r *UserRouteHandlers) RegisterUpdateUserRoute(e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	return r.RegisterUpdateUserRouteAt(r.UpdateUserPath(), e, m...)
}

// Register the handler and middleware for PUT:/user/:username at a custom path
func (r *UserRouteHandlers) RegisterUpdateUserRouteAt(path string, e EchoLike, m ...echo.MiddlewareFunc) *echo.Route {
	mw := append([]echo.MiddlewareFunc{r.UpdateUserValidator}, m...)
	return e.PUT(path, r.UpdateUserHandler, mw...)
}