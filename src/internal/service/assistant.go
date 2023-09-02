package service

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/internal/service/demo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectVersion(c *gin.Context) {
	c.String(http.StatusOK, cfg.ProjectVersion)
}

func IsAlive(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

var cfg *config.ProjectConfig

func InitService(projectConfig *config.ProjectConfig) {
	demo.InitDemo(projectConfig)
}
