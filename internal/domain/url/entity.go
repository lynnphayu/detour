package url

import (
	"detour/internal/domain/hit"
	"errors"
	"time"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
	ErrNotFound   = errors.New("URL not found")
)

type URL struct {
	ID        int        `json:"id"`
	Original  string     `json:"original"`
	Short     string     `json:"short"`
	Hits      []*hit.Hit `json:"hits"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewURL(original string) (*URL, error) {
	if original == "" {
		return nil, ErrInvalidURL
	}

	return &URL{
		Original:  original,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}
