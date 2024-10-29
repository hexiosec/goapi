package specv31

// No refs at this level as refs would be resolved inside this wrapper
type Components struct {
	Schemas         map[string]*Schema   `json:"schemas,omitempty"`
	Responses       map[string]*Response `json:"responses,omitempty"`
	Parameters      map[string]TODO      `json:"parameters,omitempty"`
	Examples        map[string]TODO      `json:"examples,omitempty"`
	RequestBodies   map[string]TODO      `json:"requestBodies,omitempty"`
	Headers         map[string]TODO      `json:"headers,omitempty"`
	SecuritySchemes map[string]TODO      `json:"securitySchemes,omitempty"`
	Links           map[string]TODO      `json:"links,omitempty"`
	Callbacks       map[string]TODO      `json:"callbacks,omitempty"`
	PathItems       map[string]TODO      `json:"pathItems,omitempty"`
}
