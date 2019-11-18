package handle

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	message := "Hello 🚕"
	messageBytes := []byte(message)
	_, err := w.Write(messageBytes)
	if err != nil {
		log.Fatalln(err)
	}
}
