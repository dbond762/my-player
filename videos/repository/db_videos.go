package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBVideosRepository interface {

}

type dbVideosRepository struct {
	db *gorm.DB
}

func NewDBVideosRepository(db *gorm.DB) *dbVideosRepository {
	return &dbVideosRepository{db}
}
