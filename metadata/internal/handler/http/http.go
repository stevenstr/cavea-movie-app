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
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	ctx := req.Context()
	m, err := h.ctrl.Get(ctx, id)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Repository get error: %v\n", err)
	}
}
