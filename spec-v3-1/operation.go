package specv31

import (
	"encoding/json"
	"slices"
)

type Operation struct {
	Tags        []string                   `json:"tags,omitempty"`
	Summary     string                     `json:"summary,omitempty"`
	Description string                     `json:"description,omitempty"`
	OperationID string                     `json:"operationId,omitempty"`
	Parameters  []*Ref[*Parameter]         `json:"parameters,omitempty"`
	RequestBody *Ref[*RequestBody]         `json:"requestBody,omitempty"`
	Responses   map[string]*Ref[*Response] `json:"responses,omitempty"`
	Deprecated  bool                       `json:"deprecated,omitempty"`
	// Security    TODO      `json:"security,omitempty"`

	Extensions map[string]any `json:"-"`
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
