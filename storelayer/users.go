package storelayer

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Handle string
	Name   string
	Posts  []Post
}

func (self *store) CreateUser(ctx context.Context, name, handle string) error {
	err := self.db.WithContext(ctx).Create(&User{
		Name:   name,
		Handle: handle,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (self *store) GetAllUsers(ctx context.Context) ([]User, error) {
	users := []User{}
	err := self.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
