// This package contains the generate commands to update the examples.
package examples

//go:generate go run ../cmd/generate -i ./petstore.yml -o ./petstore-server-go-echo/server -t server/go-echo --templates-path ../templates
//go:generate go run golang.org/x/tools/cmd/goimports@latest -w ./petstore-server-go-echo/server
