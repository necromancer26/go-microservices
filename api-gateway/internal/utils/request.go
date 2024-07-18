package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSONRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
