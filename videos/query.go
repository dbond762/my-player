package videos

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Query struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (q Query) TableName() string {
	return "queries"
}
