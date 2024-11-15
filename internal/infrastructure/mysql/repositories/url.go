package repositories

import (
	"context"
	"database/sql"
	"detour/internal/domain/url"
	"detour/internal/infrastructure/mysql/queries"
	"fmt"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepo(db *sql.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) Save(ctx context.Context, url *url.URL) (*url.URL, error) {
	result, err := r.db.ExecContext(ctx, queries.CreateURL,
		url.Original,
		url.Short,
		url.Version,
		url.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error saving URL: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	url.ID = int(id)
	return url, nil
}

func (r *URLRepository) FindMaxVersion(ctx context.Context, shortURL string) (int, error) {
	var version int
	err := r.db.QueryRowContext(ctx, queries.FindMaxVersion, shortURL).Scan(&version)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return version, nil
}

func (r *URLRepository) FindLatestByShort(ctx context.Context, shortURL string) (*url.URL, error) {
	var u url.URL
	err := r.db.QueryRowContext(ctx, queries.FindLatestURLByShort, shortURL).Scan(
		&u.ID,
		&u.Original,
		&u.Short,
		&u.Version,
		&u.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, url.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error finding URL: %w", err)
	}

	return &u, nil
}

func (r *URLRepository) CreateHit(ctx context.Context, urlID int, ip string) error {

	_, err := r.db.ExecContext(ctx, queries.CreateHit, urlID, ip)
	if err != nil {
		return fmt.Errorf("error incrementing hits: %w", err)
	}

	return nil
}

func (r *URLRepository) FindByID(ctx context.Context, id int) (*url.URL, error) {
	var u url.URL
	err := r.db.QueryRowContext(ctx, queries.FindURLByID, id).Scan(
		&u.ID,
		&u.Original,
		&u.Short,
		&u.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, url.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error finding URL: %w", err)
	}

	return &u, nil
}
