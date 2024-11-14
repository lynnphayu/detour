package queries

const (
	CreateURL = `
        INSERT INTO urls (
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
        FROM urls
        WHERE short_url = ?
    `

	FindURLByID = `
		SELECT * FROM urls WHERE id = ?
	`
)
