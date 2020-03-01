package utils

import (
	"encoding/json"
	"net/http"
)

// Decode just decodes the request body
func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
