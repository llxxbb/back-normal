package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func detMonitor(r *gin.Engine) {
	// monitor
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
