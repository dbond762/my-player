package videos

import "time"

type Video struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	PubDate     time.Time `json:"pub_date"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	Player      string    `json:"player"`
}
