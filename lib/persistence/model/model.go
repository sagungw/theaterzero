package model

import uuid "github.com/satori/go.uuid"

// Movie :nodoc:
type Movie struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	ImageURL   string `json:"image_url"`
	TrailerURL string `json:"trailer_url"`
	Synopsis   string `json:"synopsis"`
}

// Cinema :nodoc:
type Cinema struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Schedule :nodoc:
type Schedule struct {
	ID        string   `json:"id"`
	Cinema    *Cinema  `json:"cinema"`
	Movie     *Movie   `json:"movie"`
	IsPlaying bool     `json:"is_playing"`
	PlayTime  []string `json:"play_time"`
}

// GenerateID :nodoc:
func GenerateID() string {
	return uuid.NewV4().String()
}
