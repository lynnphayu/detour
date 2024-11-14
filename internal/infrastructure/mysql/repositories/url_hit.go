package repositories

import (
	"context"
	"database/sql"
	"detour/internal/domain/hit"
	"detour/internal/infrastructure/mysql/queries"
	"fmt"
)

type HitRepository struct {
	db *sql.DB
}

func NewHitRepo(db *sql.DB) *HitRepository {
	return &HitRepository{
		db: db,
	}
}

func (r *HitRepository) Save(ctx context.Context, url *hit.Hit) error {
	result, err := r.db.ExecContext(ctx, queries.CreateHit,
		url.URLID,
		url.IP,
		url.UserAgent,
		url.Referer,
		url.HitAt,
	)
	fmt.Println("error", err)

	if err != nil {
		return fmt.Errorf("error saving hit: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %w", err)
	}

	url.ID = int(id)
	fmt.Println("hit saved", url.ID)
	return nil
}

func (r *HitRepository) FindHitsByURLID(ctx context.Context, urlID int) ([]*hit.Hit, error) {
	rows, err := r.db.QueryContext(ctx, queries.FindHitsByURLID, urlID)
	if err != nil {
		return nil, fmt.Errorf("error finding hits: %w", err)
	}
	defer rows.Close()

	var hits []*hit.Hit
	for rows.Next() {
		var h hit.Hit
		if err := rows.Scan(&h.ID, &h.URLID, &h.HitAt, &h.UserAgent, &h.IP, &h.Referer, &h.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning hit: %w", err)
		}
		hits = append(hits, &h)
	}

	return hits, nil
}
