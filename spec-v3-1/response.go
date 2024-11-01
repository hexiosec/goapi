package specv31

import "encoding/json"

type Response struct {
	Description string                      `json:"description"`
	Headers     map[string]Ref[TODO]        `json:"headers,omitempty"`
	Content     map[string]*MediaTypeObject `json:"content,omitempty"`
	Links       map[string]Ref[TODO]        `json:"links,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *Response) UnmarshalJSON(value []byte) error {
	type Tintr Response

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
