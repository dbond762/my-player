package videos

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Video struct {
	ID          string `gorm:"size:20;not null" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`

	Title       string     `gorm:"size:50;not null" json:"title"`
	PubDate     time.Time  `json:"pub_date"`
	Description string     `gorm:"type:text;not null" json:"description"`
	Thumbnail   string     `gorm:"not null" json:"thumbnail"`
	Player      string     `gorm:"type:text;not null" json:"player"`
}
