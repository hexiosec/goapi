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
	}

	if !ok || target == nil {
		return fmt.Errorf("cannot de-ref %s in %s", parts[3], parts[2])
	}

	d.Value = target.(T)
	d.Ref = ""
	return nil
}
