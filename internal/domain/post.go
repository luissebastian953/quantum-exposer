package domain

import "time"

type Post struct {
	ID         int64     `json:"id"`
	Tags       []string  `json:"tags"`
	Rating     string    `json:"rating"`
	Score      int       `json:"score"`
	UploadedAt time.Time `json:"uploaded_at"`
	FileURL    string    `json:"file_url"`
}
