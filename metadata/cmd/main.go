package main

import (
	"log"
	"net/http"

	"github.com/stevenstr/cavea-movie-app/metadata/internal/controller/metadata"
	httpHandler "github.com/stevenstr/cavea-movie-app/metadata/internal/handler/http"
	"github.com/stevenstr/cavea-movie-app/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting the movie metadata service")

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httpHandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
