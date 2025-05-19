package schema

type Schema struct {
	Resources map[string]Resource `json:"resources,omitempty"`
}
