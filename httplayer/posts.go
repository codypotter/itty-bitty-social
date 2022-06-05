package httplayer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type post struct {
	Content string `json:"content"`
}

func (self *httpApi) createPost(c *gin.Context) {
	newPost := &post{}
	err := c.BindJSON(newPost)
	if err != nil {
		logrus.Error("failed to bind in create post: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// TODO: should set post.owner from auth
	err = self.app.CreatePost(c, newPost.Content, "noodle")
	if err != nil {
		logrus.Error("failed to create post: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "created successfully",
	})
}

func (self *httpApi) getAllPosts(c *gin.Context) {
	posts, err := self.app.GetAllPosts(c)
	if err != nil {
		logrus.Error("failed to get posts: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	var postList []gin.H
	for _, post := range posts {
		postList = append(postList, gin.H{
			"content": post.Content,
			"owner":   post.Owner,
		})
	}
	c.JSON(http.StatusOK, postList)
}
