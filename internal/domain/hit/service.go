package hit

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SaveHit(ctx context.Context, hit *Hit) error {
	return s.repo.Save(ctx, hit)
}

func (s *Service) GetByURLID(ctx context.Context, urlID int) ([]*Hit, error) {
	return s.repo.FindHitsByURLID(ctx, urlID)
}
