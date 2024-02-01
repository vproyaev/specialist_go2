package models

type APIResponse struct {
	Result interface{} `json:"result"`
	Err    string      `json:"error,omitempty"`
}
