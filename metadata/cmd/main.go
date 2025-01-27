package main

import (
	"log"
	"net/http"

	"movie.com/metadata/internal/controller/metadata"
	httpHandler "movie.com/metadata/internal/handler/http"
	"movie.com/metadata/internal/repository/memory"
)

func main() {
	// logger
	log.Println("Starting the movie metadata service")

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httpHandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
