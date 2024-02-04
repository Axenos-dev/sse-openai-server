package app

import (
	"fmt"
	"net/http"

	"github.com/Axenos-dev/sse-openai-server/config"
	"github.com/Axenos-dev/sse-openai-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	controllers.InitControllers(r)

	return r.Run(fmt.Sprintf(":%s", config.Config.Port))
}
