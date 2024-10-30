# goapi OpenAPI v3.1 Generator

Generate server or client implementations for OpenAPI v3.1 schemas.

## Usage

```
Usage of generate:                               
      --help                    Show help
  -i, --input string            Input schema (default "./openapi.yml")
  -o, --output string           Output folder (default "./output")
  -t, --template string         Generator template
      --templates-path string   Path to template library
  -v, --verbose                 Turn on verbose messaging
```

## Supported Languages

### Server Implementation

| Name    | Description                                                                        |
| ------- | ---------------------------------------------------------------------------------- |
| go-echo | Golang Server interface using the [Echo](https://echo.labstack.com/) web framework |

### Client Implementation

| Name     | Description                                            |
| -------- | ------------------------------------------------------ |
| ts-fetch | TypeScript Client interface using native `fetch()` API |
