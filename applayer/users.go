package applayer

import (
	"context"

	"github.com/codypotter/itty-bitty-social/storelayer"
)

func (self *app) CreateUser(ctx context.Context, name, handle string) error {
	return self.store.CreateUser(ctx, name, handle)
}

func (self *app) GetAllUsers(ctx context.Context) ([]storelayer.User, error) {
	return self.store.GetAllUsers(ctx)
}
