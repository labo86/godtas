package util

import (
	"encoding/json"
	"net/http"
)

/**
Escribe un valor en el writer como un json.
Recuerda que en caso de slices inicializarlos con make(slice, 0) para que sean  serializados como [] en vez de null.
*/
func EncodeAsJson(w http.ResponseWriter, value interface{}) {
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return
}
