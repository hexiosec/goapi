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

	for key, path := range doc.Paths {
		for method, op := range path.AsMap() {

			// If operationId is empty, generate one based on the path
			if op.OperationId == "" {
				op.OperationId = genOperationID(key, method)
				log.Warn().Msgf("Generated operationId \"%s\" for %s:%s", op.OperationId, strings.ToUpper(method), key)
			}

			// Error if operationId is not unique (inc. if generated)
			if slices.Contains(operationIDs, op.OperationId) {
				return fmt.Errorf("operation: ID %s not unique for %s:%s", op.OperationId, strings.ToUpper(method), key)
			} else {
				operationIDs = append(operationIDs, op.OperationId)
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
