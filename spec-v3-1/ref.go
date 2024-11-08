package specv31

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Ref[T any] struct {
	Ref   string `json:"$ref,omitempty"`
	Value T      `json:"value,omitempty"`
}

func (d *Ref[T]) UnmarshalJSON(value []byte) error {
	type Tintr Ref[T]

	if err := json.Unmarshal(value, (*Tintr)(d)); err != nil {
		return err
	}

	if d.Ref == "" {
		t := new(T)
		if err := json.Unmarshal(value, t); err != nil {
			return err
		}
		d.Value = *t
	}
	return nil
}

func (d Ref[T]) MarshalJSON() ([]byte, error) {
	if d.Ref != "" {
		return json.Marshal(map[string]interface{}{"$ref": d.Ref})
	}
	return json.Marshal(d.Value)
}

func (d *Ref[T]) DeRef(lookup *Components) error {
	if d.Ref == "" {
		return nil
	}

	parts := strings.Split(d.Ref, "/")

	if len(parts) != 4 || parts[0] != "#" || parts[1] != "components" {
		return fmt.Errorf("unsupported $ref format: %s", d.Ref)
	}

	var target any
	ok := false

	switch parts[2] {
	case "schemas":
		target, ok = lookup.Schemas[parts[3]]
	case "parameters":
		target, ok = lookup.Parameters[parts[3]]
	case "responses":
		target, ok = lookup.Responses[parts[3]]
	case "requestBodies":
		target, ok = lookup.RequestBodies[parts[3]]
	case "headers":
		target, ok = lookup.Headers[parts[3]]
	case "securitySchemes":
		target, ok = lookup.SecuritySchemes[parts[3]]
	}

	if !ok || target == nil {
		return fmt.Errorf("cannot de-ref %s in %s", parts[3], parts[2])
	}

	if val, ok := target.(T); ok {
		d.Value = val
	} else {
		return fmt.Errorf("found %s in %s but incorrect type returned for ref: %v", parts[3], parts[2], target)
	}

	d.Ref = ""
	return nil
}
