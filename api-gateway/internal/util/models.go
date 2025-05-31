package util

type ErrorResponse struct {
	Err  []string `json:"errors"`
	Data any      `json:"data,omitempty"`
}
