package specv31

import (
	"encoding/json"
	"strings"
)

type Extensible struct {
	Extensions map[string]any `json:"-,omitempty"`
}

func HandleExtensions(value []byte) (map[string]any, error) {
	raw := map[string]any{}

	if err := json.Unmarshal(value, &raw); err != nil {
		return nil, err
	}

	for key := range raw {
		if !strings.HasPrefix(key, "x-") {
			delete(raw, key)
		}
	}

	if len(raw) > 0 {
		return raw, nil
	}

	return nil, nil
}
