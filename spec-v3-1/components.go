package specv31

import "strings"

// No refs at this level as refs would be resolved inside this wrapper
type Components struct {
	Schemas         map[string]*Schema         `json:"schemas,omitempty"`
	Responses       map[string]*Response       `json:"responses,omitempty"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty"`
	Examples        map[string]TODO            `json:"examples,omitempty"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]TODO            `json:"headers,omitempty"`
	SecuritySchemes map[string]TODO            `json:"securitySchemes,omitempty"`
	Links           map[string]TODO            `json:"links,omitempty"`
	Callbacks       map[string]TODO            `json:"callbacks,omitempty"`
	PathItems       map[string]*PathItemObject `json:"pathItems,omitempty"`
}

func (d *Components) GetParameter(ref string) *Parameter {
	key := strings.TrimPrefix("#/components/parameters/", ref)
	if p, ok := d.Parameters[key]; ok {
		return p
	}
	return nil
}
