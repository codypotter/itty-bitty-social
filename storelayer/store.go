package storelayer

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store interface {
	CreateUser(ctx context.Context, name, handle string) error
	GetAllUsers(ctx context.Context) ([]User, error)
	CreatePost(ctx context.Context, content, owner string) error
	GetAllPosts(ctx context.Context) ([]Post, error)
}

type store struct {
	db *gorm.DB
}

func New() *store {
	db, err := gorm.Open(sqlite.Open("itty-bitty-social.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	return &store{
		db: db,
	}
}
