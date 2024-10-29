package specv31

import "encoding/json"

type TODO map[string]any

type Document struct {
	OpenAPI      string                    `json:"openapi"`
	Info         *Info                     `json:"info"`
	Schema       string                    `json:"$schema,omitempty"`
	Servers      []*Server                 `json:"servers,omitempty"`
	Paths        map[string]PathItemObject `json:"paths,omitempty"`
	Webhooks     map[string]TODO           `json:"webhooks,omitempty"`
	Components   *Components               `json:"components,omitempty"`
	Security     []TODO                    `json:"security,omitempty"`
	Tags         []*Tag                    `json:"tags,omitempty"`
	ExternalDocs TODO                      `json:"externalDocs,omitempty"`

	Extensions map[string]any `json:"-,omitempty"`
}

func (d *Document) UnmarshalJSON(value []byte) error {
	type Tintr Document

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
