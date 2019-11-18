package handle

import (
	"encoding/json"
	"log"
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
	response := map[string]string{
		"error": err.Error(),
	}
	respondWithJSON(w, response, http.StatusInternalServerError)
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
