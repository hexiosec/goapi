package specv31

import "encoding/json"

type RequestBody struct {
	Description string                      `json:"description,omitempty"`
	Content     map[string]*MediaTypeObject `json:"content,omitempty"`
	Required    bool                        `json:"required,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *RequestBody) UnmarshalJSON(value []byte) error {
	type Tintr RequestBody

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

type MediaTypeObject struct {
	Schema   *Ref[*Schema]   `json:"schema,omitempty"`
	Encoding map[string]TODO `json:"encoding,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *MediaTypeObject) UnmarshalJSON(value []byte) error {
	type Tintr MediaTypeObject

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
