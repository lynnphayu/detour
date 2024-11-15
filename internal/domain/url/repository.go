package url

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, url *URL) (*URL, error)
	FindLatestByShort(ctx context.Context, short string) (*URL, error)
	FindByID(ctx context.Context, id int) (*URL, error)
	CreateHit(ctx context.Context, urlID int, ip string) error
	FindMaxVersion(ctx context.Context, originalURL string) (int, error)
}
