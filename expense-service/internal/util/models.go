package util

type ErrorJSONRespone struct {
	Error bool              `json:"error"`
	Msgs  map[string]string `json:"error-messages,omitempty"`
}
