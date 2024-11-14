package hit

import "context"

type Repository interface {
	Save(ctx context.Context, hit *Hit) error
	FindHitsByURLID(ctx context.Context, urlID int) ([]*Hit, error)
}
