package schema

type HelmBody struct {
	Chart     map[string]string      `json:"chart_url"`
	Name      string                 `json:"name"`
	NameSpace string                 `json:"namespace"`
	Values    map[string]interface{} `json:"values"`
}

type Error struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"description,omitempty"`
}
