package specv31

import "encoding/json"

type Server struct {
	URL         string                           `json:"url"`
	Description string                           `json:"description,omitempty"`
	Variables   map[string]*ServerVariableObject `json:"variables,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *Server) UnmarshalJSON(value []byte) error {
	type Tintr Server

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

type ServerVariableObject struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default"`
	Description string   `json:"description,omitempty"`

	Extensions map[string]any `json:"-"`
}

func (d *ServerVariableObject) UnmarshalJSON(value []byte) error {
	type Tintr ServerVariableObject

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
