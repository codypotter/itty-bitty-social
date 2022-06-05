package applayer

import (
	"context"
)

type Post struct {
	Owner   string
	Content string
}

func (self *app) CreatePost(ctx context.Context, content, owner string) error {
	return self.store.CreatePost(ctx, content, owner)
}

func (self *app) GetAllPosts(ctx context.Context) ([]Post, error) {
	postRows, err := self.store.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for _, postRow := range postRows {
		posts = append(posts, Post{
			Owner:   postRow.User.Handle,
			Content: postRow.Content,
		})
	}
	return posts, nil
}
