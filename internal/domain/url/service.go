package url

import (
	"context"
	"crypto/rand"
	"encoding/base64"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateShortURL(ctx context.Context, originalURL string) (*URL, error) {
	url, err := NewURL(originalURL)
	if err != nil {
		return nil, err
	}

	url.Short = generateShortURL()

	if err := s.repo.Save(ctx, url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *Service) GetByShortURL(ctx context.Context, shortURL string) (*URL, error) {
	return s.repo.FindByShort(ctx, shortURL)
}

func generateShortURL() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}

func (s *Service) IncrementHits(ctx context.Context, urlID string, ip string) error {
	return s.repo.CreateHit(ctx, urlID, ip)
}
