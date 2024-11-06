package specv31

import (
	"encoding/json"
	"slices"
)

type PathItemObject struct {
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	Options     *Operation `json:"options,omitempty"`
	Head        *Operation `json:"head,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Trace       *Operation `json:"trace,omitempty"`
}

func (d PathItemObject) AsMap() map[string]*Operation {
	m := map[string]*Operation{}

	if d.Get != nil {
		m["get"] = d.Get
	}
	if d.Put != nil {
		m["put"] = d.Put
	}
	if d.Post != nil {
		m["post"] = d.Post
	}
	if d.Delete != nil {
		m["delete"] = d.Delete
	}
	if d.Options != nil {
		m["options"] = d.Options
	}
	if d.Head != nil {
		m["head"] = d.Head
	}
	if d.Patch != nil {
		m["patch"] = d.Patch
	}
	if d.Trace != nil {
		m["trace"] = d.Trace
	}

	return m
}

type Operation struct {
	Tags        []string          `json:"tags,omitempty"`
	Summary     string            `json:"summary,omitempty"`
	Description string            `json:"description,omitempty"`
	OperationId string            `json:"operationId,omitempty"`
	Parameters  []Ref[*Parameter] `json:"parameters,omitempty"`
	// RequestBody Ref[TODO] `json:"requestBody,omitempty"`
	Responses  TODO `json:"responses,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
	// Security    TODO      `json:"security,omitempty"`

	Extensions map[string]any `json:"-,omitempty"`
}

func (d *Operation) UnmarshalJSON(value []byte) error {
	type Tintr Operation

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

func (d *Operation) HasTag(name string) bool {
	return slices.Contains(d.Tags, name)
}

type Parameter struct {
	Name            string       `json:"name"`
	In              string       `json:"in"`
	Description     string       `json:"description,omitempty"`
	Required        bool         `json:"required,omitempty"`
	Deprecated      bool         `json:"deprecated,omitempty"`
	AllowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	Schema          Ref[*Schema] `json:"schema,omitempty"`
}
