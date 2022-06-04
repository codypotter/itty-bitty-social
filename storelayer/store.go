package storelayer

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Layer interface{}

type store struct {
	db *gorm.DB
}

func New() *store {
	db, err := gorm.Open(sqlite.Open("itty-bitty-social.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &store{
		db: db,
	}
}
