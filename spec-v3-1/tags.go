package specv31

import "encoding/json"

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *Tag) UnmarshalJSON(value []byte) error {
	type Tintr Tag

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
