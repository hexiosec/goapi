package generator

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ettle/strcase"
	"github.com/rs/zerolog/log"

	specv31 "github.com/hexiosec/goapi/spec-v3-1"
)

func Validate(doc *specv31.Document) error {
	operationIDs := []string{}
	defaultTagAdded := false

	for key, path := range doc.Paths {
		for method, op := range path.AsMap() {

			// If operationId is empty, generate one based on the path
			if op.OperationID == "" {
				op.OperationID = genOperationID(key, method)
				log.Warn().Msgf("Generated operationId \"%s\" for %s:%s", op.OperationID, strings.ToUpper(method), key)
			}

			// Error if operationId is not unique (inc. if generated)
			if slices.Contains(operationIDs, op.OperationID) {
				return fmt.Errorf("operation: ID %s not unique for %s:%s", op.OperationID, strings.ToUpper(method), key)
			} else {
				operationIDs = append(operationIDs, op.OperationID)
			}

			// If no tags specified, add "default"
			if len(op.Tags) == 0 {
				op.Tags = []string{"default"}
				if !defaultTagAdded {
					doc.Tags = append(doc.Tags, &specv31.Tag{
						Name:        "default",
						Description: "Default tag for untagged routes",
					})
					defaultTagAdded = true
				}
			}

			// If requestBody is set and not a ref, move to schemas
			if op.RequestBody != nil {
				if op.RequestBody.Value != nil {
					for mime, mto := range op.RequestBody.Value.Content {
						if mime == "application/json" {
							if mto.Schema != nil {
								if mto.Schema.Value != nil {
									key := strcase.ToGoPascal(op.OperationID) + "JSONRequest"
									log.Debug().Msgf("Moving MediaTypeObject schema from operation %s requestBody to %s", op.OperationID, key)
									doc.Components.Schemas[key] = mto.Schema.Value
									mto.Schema.Ref = "#/components/schemas/" + key
									mto.Schema.Value = nil
									op.Extensions["x-goapi-json-requestbody-object"] = key
								}
							}
						}
					}
				}
			}

			// If a JSON response is present and not a ref, move to schemas
			for status, ref := range op.Responses {
				if ref.Value != nil {
					for mime, mto := range ref.Value.Content {
						if mime == "application/json" {
							if mto.Schema != nil {
								if mto.Schema.Value != nil {
									key := strcase.ToGoPascal(op.OperationID) + "JSON" + status + "Response"
									log.Debug().Msgf("Moving MediaTypeObject schema from operation %s %s response to %s", op.OperationID, status, key)
									doc.Components.Schemas[key] = mto.Schema.Value
									mto.Schema.Ref = "#/components/schemas/" + key
									mto.Schema.Value = nil
								}
							}
						}
					}
				}
			}

			// DeRef params
			for _, param := range op.Parameters {
				if err := param.DeRef(doc.Components); err != nil {
					return err
				}
				if err := param.Value.Schema.DeRef(doc.Components); err != nil {
					return err
				}
			}

			// Create Query object if needed
			qry := &specv31.Schema{
				Type:       "object",
				Properties: map[string]specv31.Ref[*specv31.Schema]{},
				Extensions: map[string]any{
					"x-goapi-query-schema": true,
				},
			}
			for _, ref := range op.Parameters {
				param := ref.Value
				if param.In == "query" {
					qry.Properties[param.Name] = *param.Schema
					if param.Required {
						qry.Required = append(qry.Required, param.Name)
					}
				}
			}
			if len(qry.Properties) > 0 {
				key := strcase.ToGoPascal(op.OperationID) + "Query"
				log.Debug().Msgf("Creating new Query schema from operation %s as %s", op.OperationID, key)
				doc.Components.Schemas[key] = qry
				op.Extensions["x-goapi-query-object"] = key
			}
		}
	}

	return nil
}

func genOperationID(path string, method string) string {
	parts := []string{method}

	for _, str := range strings.Split(strings.Trim(path, "/"), "/") {
		if str != "" {
			if strings.HasPrefix(str, "{") {
				parts = append(parts, "by_"+strings.Trim(str, "{}"))
			} else {
				parts = append(parts, str)
			}
		}
	}

	if len(parts) == 1 {
		parts = append(parts, "index")
	}

	return strcase.ToCamel(strings.Join(parts, "_"))
}
