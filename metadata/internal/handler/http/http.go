package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/stevenstr/cavea-movie-app/metadata/internal/controller/metadata"
	"github.com/stevenstr/cavea-movie-app/metadata/internal/repository"
)

type Handler struct {
	ctrl *metadata.Controller
}

func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	m, err := h.ctrl.Get(req.Context(), id)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("Repository get error: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Repository get error: %v\n", err)
	}
}
