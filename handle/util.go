package handle

import (
	"encoding/json"
	"log"
	"meme/types"
	"net/http"
)

func readBody(r *http.Request, dest interface{}) error {
	return json.NewDecoder(r.Body).Decode(dest)
}

func respondWithStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func respondWithError(w http.ResponseWriter, err error) {
	log.Println(err)

	status := http.StatusInternalServerError
	message := "an internal error occurred"
	if se, ok := err.(types.StatusError); ok {
		status = se.Status
		message = se.Message
	}

	response := map[string]string{
		"error": message,
	}
	respondWithJSON(w, response, status)
}

func respondWithJSON(w http.ResponseWriter, value interface{}, status int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
