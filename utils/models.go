package models

import (
	"time"
)

type Video struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnails  string    `json:"thumbnails"`
	PublishedAt time.Time `json:"published_at"`
}
