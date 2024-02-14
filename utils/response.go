package utils

import (
	"encoding/json"
	"net/http"
	"log"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	log.Println(message)
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	log.Println(string(response))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
