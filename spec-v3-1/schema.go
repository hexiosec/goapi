package specv31

import "encoding/json"

// Subset of supported Schema fields only
type Schema struct {
	AllOf []Ref[*Schema] `json:"allOf,omitempty"`
	AnyOf []Ref[*Schema] `json:"anyOf,omitempty"`
	OneOf []Ref[*Schema] `json:"oneOf,omitempty"`

	// Arrays
	Items Ref[*Schema] `json:"items,omitempty"`

	// Objects
	Properties           map[string]Ref[*Schema] `json:"properties,omitempty"`
	AdditionalProperties Ref[*Schema]            `json:"additionalProperties,omitempty"`

	// Strings
	MaxLength int    `json:"maxLength,omitempty"`
	MinLength int    `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`

	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	Required    []string `json:"required,omitempty"`
	Enum        []any    `json:"enum,omitempty"`

	Extensions map[string]any `json:"-,omitempty"`
}

func (d *Schema) UnmarshalJSON(value []byte) error {
	type Tintr Schema

	if err := json.Unmarshal(value, (*Tintr)(d)); err != nil {
		return err
	}

	if ext, err := HandleExtensions(value); err == nil {
		d.Extensions = ext
	} else {
		return err
	}

	return nil
}
