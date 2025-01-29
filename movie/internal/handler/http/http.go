package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"movie.com/movie/internal/controller/movie"
)

type Handler struct {
	ctrl *movie.Controller
}

func New(ctrl *movie.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.Get(req.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}

	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
