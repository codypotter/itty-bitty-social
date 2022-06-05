package httplayer

import (
	"net/http"

	"github.com/codypotter/itty-bitty-social/applayer"
	"github.com/gin-gonic/gin"
)

type httpApi struct {
	engine *gin.Engine
	app    applayer.App
}

func New(appLayer applayer.App) *httpApi {
	a := &httpApi{
		engine: gin.New(),
		app:    appLayer,
	}
	a.SetupRoutes()
	return a
}

func (self *httpApi) SetupRoutes() {
	self.engine.Use(gin.Recovery())
	api := self.engine.Group("/api")
	{
		api.GET("/ping", pong)
		users := api.Group("/users")
		{
			users.GET("", self.getAllUsers)
			users.POST("", self.createUser)
		}
		posts := api.Group("/posts")
		{
			posts.GET("", self.getAllPosts)
			posts.POST("", self.createPost)
		}
	}
}

func (self *httpApi) Engage() {
	self.engine.Run()
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
