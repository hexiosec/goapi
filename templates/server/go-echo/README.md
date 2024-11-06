# Go Echo Server

Golang Echo-compatible server interface wrapper

## Supported Features

- Typed server stub wrapper for path operations
- Typed model outputs for schemas

## Limitations

- Only one successful JSON response body type per operation can be handled by the wrapper

## Extension Parameters

Output can be customised by applying extension parameters to the relevant parts of the OpenAPI spec.

| Parameter                    | Scope  | Description                                         |
| ---------------------------- | ------ | --------------------------------------------------- |
| `x-go-name`                  | Schema | Override struct field name for a given child schema |
| `x-go-type`                  | Schema | Override the type name for a given schema           |
| `x-go-skip-optional-pointer` | Schema | Don't add pointer for optional field                |
| `x-validate-extra-tags`      | Schema | Extra tags to add to the `validate` tag             |
| `x-validate-override-tags`   | Schema | Override the `validate` tag                         |
