package url

import (
	"errors"
	"time"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
	ErrNotFound   = errors.New("URL not found")
)

type URL struct {
	ID        string
	Original  string
	Short     string
	CreatedAt time.Time
	UpdatedAt time.Time
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
