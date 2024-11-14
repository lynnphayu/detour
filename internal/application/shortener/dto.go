package shortener

import "time"

type CreateURLDTO struct {
	OriginalURL string `json:"url" validate:"required,url"`
}

type URLResponseDTO struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type HeaderDTO struct {
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Referer   string `json:"referer"`
}
