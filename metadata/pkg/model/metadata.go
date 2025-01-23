package model

// Metadata defines the movie metadata.
type Metadata struct {
	ID          string `json:"id"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
	Cast        string `json:"cast"`
	Script      string `json:"script"`
	Premiere    string `json:"premiere"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
}
