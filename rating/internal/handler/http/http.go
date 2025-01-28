package http

import (
	"net/http"

	"movie.com/rating/internal/controller/rating"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func HandleGetPutRating(w http.ResponseWriter, req *http.Request) {

}
