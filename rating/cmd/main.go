package main

import (
	"log"
	"net/http"

	"github.com/stevenstr/cavea-movie-app/rating/internal/controller/rating"
	httpHandler "github.com/stevenstr/cavea-movie-app/rating/internal/handler/http"
	"github.com/stevenstr/cavea-movie-app/rating/internal/repository/memory"
)

func main() {
	log.Println("Staring the movie rating service")

	mux := http.NewServeMux()
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)
	mux.HandleFunc("/rating", h.HandleRating)

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
