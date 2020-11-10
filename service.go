package godtas

import (
	"encoding/json"
	"net/http"
)

/**
Escribe un valor en el writer como un json
*/
func EncodeAsJson(w http.ResponseWriter, value interface{}) {
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return
}
