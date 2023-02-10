package api

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/internal/service"
	"cdel/demo/Normal/internal/service/demo"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	if config.CTX.C.GinRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.SetTrustedProxies(nil)

	// 用于监控
	r.GET("/monitorDB/monitor", service.IsAlive)
	r.GET("/monitorDB/monitor.shtml", service.IsAlive)
	r.GET("isAlive", service.IsAlive)

	// 项目版本查询
	r.GET("/version", service.ProjectVersion)

	gDemo := r.Group("/demo")
	{
		gDemo.POST("/v1", demo.V1)
		gDemo.POST("/v2", demo.V2)
	}
	return r
}
