package queries

const (
	CreateURL = `
        INSERT INTO url (
            original_url,
            short_url,
            created_at,
            updated_at
        ) VALUES (?, ?, ?, ?)
    `

	FindURLByShort = `
        SELECT 
            id,
            original_url,
            short_url,
            created_at,
            updated_at
        FROM url 
        WHERE short_url = ?
    `

	CreateHit = `
        INSERT INTO hit (
            url_id, 
            ip,
            hit_at
        ) VALUES (?, ?, NOW())
    `

	FindURLByID = `
		SELECT * FROM url WHERE id = ?
	`
)
