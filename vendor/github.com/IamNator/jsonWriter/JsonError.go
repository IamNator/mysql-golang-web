package jsonWriter

import (
	"encoding/json"
	"net/http"
)

//Writes Error Message to ResponseWriter in json format
//Similar to http.Error(w ResponseWriter, error string, code int)
//Example of json written is --> {"error":ErrorMessage}
func Error(w http.ResponseWriter, ErrorMessage string, ErrorCode int) {
	w.WriteHeader(ErrorCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: ErrorMessage,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to encode json response"))
	}
}

