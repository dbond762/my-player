package user

import (
	"github.com/dbond762/my-player/videos"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name     string         `gorm:"size:50;unique;not null" json:"name"`
	Password string         `gorm:"size:60;not null" json:"-"`
	Queries  []videos.Query `gorm:"many2many:user_queries"`
}
