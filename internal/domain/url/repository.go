package url

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, url *URL) error
	FindByShort(ctx context.Context, short string) (*URL, error)
	FindByID(ctx context.Context, id string) (*URL, error)
	CreateHit(ctx context.Context, urlID string, ip string) error
}
