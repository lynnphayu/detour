package queries

const (
	CreateHit = `
        INSERT INTO url_hits (
            url_id,
            ip,
            user_agent,
            referer,
            hit_at,
            created_at
        ) VALUES (?, ?, ?, ?, ?, NOW())
    `

	FindHitsByURLID = `
		SELECT * FROM url_hits WHERE url_hits.url_id = ?
	`
)
