package specv31

type PathItemObject struct {
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	Options     *Operation `json:"options,omitempty"`
	Head        *Operation `json:"head,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Trace       *Operation `json:"trace,omitempty"`
}

func (d PathItemObject) AsMap() map[string]*Operation {
	m := map[string]*Operation{}

	if d.Get != nil {
		m["get"] = d.Get
	}
	if d.Put != nil {
		m["put"] = d.Put
	}
	if d.Post != nil {
		m["post"] = d.Post
	}
	if d.Delete != nil {
		m["delete"] = d.Delete
	}
	if d.Options != nil {
		m["options"] = d.Options
	}
	if d.Head != nil {
		m["head"] = d.Head
	}
	if d.Patch != nil {
		m["patch"] = d.Patch
	}
	if d.Trace != nil {
		m["trace"] = d.Trace
	}

	return m
}
