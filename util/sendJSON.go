package util

import (
	"encoding/json"
	"net/http"
)

// SendJSON format the http response in JSON
func SendJSON(w http.ResponseWriter, status int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}
