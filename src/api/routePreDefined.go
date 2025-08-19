package api

import (
	"back/demo/internal/service"

	"github.com/gin-gonic/gin"
)

func RoutePreDefined(r *gin.Engine) {
	// 用于监控
	r.GET("/monitorDB/monitor", service.IsAlive)
	r.GET("/monitorDB/monitor.shtml", service.IsAlive)
	r.GET("isAlive", service.IsAlive)

	// 项目版本查询
	r.GET("/version", service.ProjectVersion)
}
