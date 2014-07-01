package maxcdn

import (
	"encoding/json"
	"net/http"
)

// Response object for all json requests.
type Response struct {
	Code  int             `json:"code,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"`
	Error struct {
		Message string `json:"message,omitempty"`
		Type    string `json:"type,omitempty"`
	} `json:"error,omitempty"`

	// Non-JSON data.
	Headers *http.Header
}

// Generic is the generic data type for JSON responses from API calls.
type Generic map[string]interface{}
