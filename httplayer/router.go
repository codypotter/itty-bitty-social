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
		engine: gin.Default(),
		app:    appLayer,
	}
	a.SetupRoutes()
	return a
}

func (a *httpApi) SetupRoutes() {
	api := a.engine.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}

func (a *httpApi) Engage() {
	a.engine.Run()
}
