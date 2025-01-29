package main

import (
	"log"
	"net/http"

	"movie.com/rating/internal/controller/rating"
	httpHandler "movie.com/rating/internal/handler/http"
	"movie.com/rating/internal/repository/memory"
)

func main() {
	log.Println("Staring the movie rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httpHandler.New(ctrl)

	http.Handle("/rating", http.HandlerFunc(h.HandleRating))

	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
