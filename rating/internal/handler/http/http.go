package http

import (
	"net/http"

	"movie.com/rating/internal/controller/rating"
	"movie.com/rating/pkg/model"
)

// Handler defines rating service controller.
type Handler struct {
	ctrl *rating.Controller
}

// New creates a new rating service http handler.
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// HandleRating handles Get and Put methods by /rating uri requests.
func HandleRating(w http.ResponseWriter, req *http.Request) {
	recordID := model.RecordID(req.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	recordType := model.RecordType(req.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
	case http.MethodPut:
	default:
		w.WriteHeader(http.StatusBadRequest) // 400
	}
}
