package main

import (
	"log"
	"os"

	"github.com/mikitachab/score-board/server"
)

func main() {
	s := server.NewServer()
	port := getPort()
	log.Printf("Starting server on port%s ...\n", port)
	log.Fatal(s.ListenAndServe(port))
}

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}
