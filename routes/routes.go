package routes

import (
	"gin-web-app/logger"
	"gin-web-app/settings"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(settings.Conf.Mode)
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	return r
}
