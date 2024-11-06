package specv31

import (
	"encoding/json"
)

type Info struct {
	Title          string   `json:"title"`
	Summary        string   `json:"summary,omitempty"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"terms_of_service,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        TODO     `json:"license,omitempty"`
	Version        string   `json:"version"`

	Extensions map[string]any `json:"-,omitempty"`
}

func (d *Info) UnmarshalJSON(value []byte) error {
	type Tintr Info

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

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`

	Extensions map[string]any `json:"-,omitempty"`
}

func (d *Contact) UnmarshalJSON(value []byte) error {
	type Tintr Contact

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

type License struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`

	Extensions map[string]any `json:"-,omitempty"`
}

func (d *License) UnmarshalJSON(value []byte) error {
	type Tintr License

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
