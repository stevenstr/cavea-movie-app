package model

import "movie.com/metadata/pkg/model"

// MovieDetails include movie metadata
// and its aggregated reting.
type MovieDetails struct {
	Rating   float64        `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
