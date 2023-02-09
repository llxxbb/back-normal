package service

import (
	"cdel/demo/Normal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectVersion(c *gin.Context) {
	c.String(http.StatusOK, config.C.ProjectVersion)
}

func IsAlive(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
