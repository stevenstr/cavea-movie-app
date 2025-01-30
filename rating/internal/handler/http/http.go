package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/stevenstr/cavea-movie-app/rating/internal/controller/rating"
	"github.com/stevenstr/cavea-movie-app/rating/pkg/model"
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
func (h *Handler) HandleRating(w http.ResponseWriter, req *http.Request) {
	recordID := model.RecordID(req.FormValue("id"))
	if recordID == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(req.FormValue("type"))
	if recordType == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(req.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}

	case http.MethodPut:
		userID := model.UserID(req.FormValue("userId"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(req.Context(), recordID, recordType, &model.Rating{UserID: userID, Value: model.RatingValue(v)}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
