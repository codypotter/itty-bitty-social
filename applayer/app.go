package applayer

import (
	"context"

	"github.com/codypotter/itty-bitty-social/storelayer"
)

type App interface {
	GetAllUsers(ctx context.Context) ([]storelayer.User, error)
	CreateUser(ctx context.Context, name, handle string) error
	CreatePost(ctx context.Context, content, owner string) error
	GetAllPosts(ctx context.Context) ([]Post, error)
}

type app struct {
	store storelayer.Store
}

func New(store storelayer.Store) *app {
	return &app{
		store: store,
	}
}
