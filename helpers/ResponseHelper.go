package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if status == 500 && data == nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(data)
}
