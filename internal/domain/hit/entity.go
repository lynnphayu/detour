package hit

import (
	"time"
)

type Hit struct {
	ID        int       `json:"id"`
	URLID     int       `json:"url_id"`
	HitAt     time.Time `json:"hit_at"`
	UserAgent string    `json:"user_agent"`
	IP        string    `json:"ip"`
	Referer   string    `json:"referer"`
	CreatedAt time.Time `json:"created_at"`
}
