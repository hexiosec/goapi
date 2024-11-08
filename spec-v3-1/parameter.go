package specv31

import "encoding/json"

type Parameter struct {
	Name            string        `json:"name"`
	In              string        `json:"in"`
	Description     string        `json:"description,omitempty"`
	Required        bool          `json:"required,omitempty"`
	Deprecated      bool          `json:"deprecated,omitempty"`
	AllowEmptyValue bool          `json:"allowEmptyValue,omitempty"`
	Schema          *Ref[*Schema] `json:"schema,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *Parameter) UnmarshalJSON(value []byte) error {
	type Tintr Parameter

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
