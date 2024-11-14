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

func (r *URLRepository) Save(ctx context.Context, url *url.URL) error {
	result, err := r.db.ExecContext(ctx, queries.CreateURL,
		url.Original,
		url.Short,
		url.CreatedAt,
		url.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error saving URL: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %w", err)
	}

	url.ID = fmt.Sprintf("%d", id)
	return nil
}

func (r *URLRepository) FindByShort(ctx context.Context, shortURL string) (*url.URL, error) {
	var u url.URL
	err := r.db.QueryRowContext(ctx, queries.FindURLByShort, shortURL).Scan(
		&u.ID,
		&u.Original,
		&u.Short,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, url.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error finding URL: %w", err)
	}

	return &u, nil
}

func (r *URLRepository) CreateHit(ctx context.Context, urlID string, ip string) error {

	_, err := r.db.ExecContext(ctx, queries.CreateHit, urlID, ip)
	if err != nil {
		return fmt.Errorf("error incrementing hits: %w", err)
	}

	return nil
}

func (r *URLRepository) FindByID(ctx context.Context, id string) (*url.URL, error) {
	var u url.URL
	err := r.db.QueryRowContext(ctx, queries.FindURLByID, id).Scan(
		&u.ID,
		&u.Original,
		&u.Short,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, url.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error finding URL: %w", err)
	}

	return &u, nil
}
