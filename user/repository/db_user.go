package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBUserRepository interface {

}

type dbUserRepository struct {
	db *gorm.DB
}

func NewDBUserRepository(db *gorm.DB) *dbUserRepository {
	return &dbUserRepository{db}
}
