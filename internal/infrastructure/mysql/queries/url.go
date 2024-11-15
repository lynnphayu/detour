package queries

const (
	CreateURL = `
        INSERT INTO urls (
            original_url,
            short_url,
            version,
            created_at
        ) VALUES (?, ?, ?, ?)
    `

	FindLatestURLByShort = `
        SELECT 
            id,
            original_url,
            short_url,
            version,
            created_at
        FROM urls
        WHERE short_url = ?
        ORDER BY version DESC
        LIMIT 1
    `

	FindURLByID = `
		SELECT * FROM urls WHERE id = ?
	`

	FindMaxVersion = `
		SELECT MAX(version) FROM urls WHERE short_url = ?
	`
)
