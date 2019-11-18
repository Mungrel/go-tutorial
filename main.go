package main

import (
	"log"
	"meme/server"
	"net/http"
)

func main() {
	log.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", server.New())
	if err != nil {
		log.Fatalln(err)
	}
}
