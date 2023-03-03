package api

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/internal/service"
	"cdel/demo/Normal/internal/service/demo"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ppgin "github.com/pinpoint-apm/pinpoint-go-agent/plugin/gin"
	"go.uber.org/zap"
)

func SetupRouter() *gin.Engine {
	cfg := config.CTX.Cfg
	if cfg.GinRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.SetTrustedProxies(nil)

	// 设置中间间-----------------------------------
	// combine with zap
	//r.Use(ginzap.Ginzap(zap.L(), tool.LogTmFmtWithMS, false))		// 可以打印所有的请求日志
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// use pinpoint
	cfg.PinPoint.InitPinPoint(cfg.ProjectName, cfg.Host)
	r.Use(ppgin.Middleware())

	// 设置中间间-----------------------------------
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
		gDemo.POST("/rest", demo.RemoteCall)
	}
	gTmp := r.Group("/tmp")
	{
		gTmp.POST("/byName", demo.DbSelect)
		gTmp.POST("/byNameCached", demo.DbSelectCached)
		gTmp.POST("/timeout", demo.DBTimeout)
	}
	gPinPoint := r.Group("/pp")
	{
		gPinPoint.GET("/app2", demo.App2)
	}
	return r
}
