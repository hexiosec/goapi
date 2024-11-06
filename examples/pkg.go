// This package contains the generate commands to update the examples.
package examples

//go:generate go run .. -i ./petstore.yml -o ./petstore-server-go-echo/server -t server/spec --clean
//go:generate go run .. -i ./petstore.yml -o ./petstore-server-go-echo/server -t server/go-echo
//go:generate go run golang.org/x/tools/cmd/goimports@latest -w ./petstore-server-go-echo/server

//go:generate go run .. -i ./petstore3.yml -o ./petstore3-server-go-echo/server -t server/spec --clean
//go:generate go run .. -i ./petstore3.yml -o ./petstore3-server-go-echo/server -t server/go-echo
//go:generate go run golang.org/x/tools/cmd/goimports@latest -w ./petstore3-server-go-echo/server
