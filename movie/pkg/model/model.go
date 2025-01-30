package model

import "github.com/stevenstr/cavea-movie-app/metadata/pkg/model"

// MovieDetails include movie metadata
// and its aggregated reting.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
