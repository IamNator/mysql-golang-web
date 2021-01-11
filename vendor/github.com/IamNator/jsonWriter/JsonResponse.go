package jsonWriter

import (
	"encoding/json"
	"net/http"
)

//Writes a Response Message to ResponseWriter in json format
//Similar to http.Write(string)
//Example of json written is --> {"title": Title, "message":Message}
func Response(w http.ResponseWriter, Title string, Message string, StatusCode int) {
	w.WriteHeader(StatusCode)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(struct {
		Title string  `json:"title"`
		Message string `json:"message"`
	}{
		Title: Title,
		Message: Message,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to encode json response"))
	}
}



