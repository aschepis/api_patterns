package rendering

import (
	"encoding/json"
	"net/http"
)

type RenderFunc func(w http.ResponseWriter, response interface{}) error

func RenderJSON(w http.ResponseWriter, response interface{}) error {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	return encoder.Encode(response)
}
