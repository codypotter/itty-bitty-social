package storelayer

import (
	"context"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	Content string
	User    User
}

func (self *store) GetAllPosts(ctx context.Context) ([]Post, error) {
	posts := []Post{}
	err := self.db.WithContext(ctx).Preload("User").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (self *store) CreatePost(ctx context.Context, content, owner string) error {
	user := &User{}
	err := self.db.Limit(1).Find(user, "handle = ?", owner).Error
	if err != nil {
		return err
	}
	err = self.db.WithContext(ctx).Create(&Post{
		Content: content,
		UserID:  user.ID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
