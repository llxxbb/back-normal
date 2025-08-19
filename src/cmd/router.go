package main

import (
	"back/demo/api"
	"back/demo/config"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	ppgin "github.com/pinpoint-apm/pinpoint-go-agent/plugin/gin"
	"go.uber.org/zap"
)

func SetupRouter(cfg *config.ProjectConfig) *gin.Engine {
	if cfg.GinRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	_ = r.SetTrustedProxies(nil)
	r.UseH2C = true

	// 设置中间间-----------------------------------
	// combine with zap
	//r.Use(ginzap.Ginzap(zap.L(), tool.LogTmFmtWithMS, false))		// 可以打印所有的请求日志
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// use pinpoint
	r.Use(ppgin.Middleware())

	// 路由信息 -----------------------------------------------
	// 需要自行定制的路由
	api.RouteApp(r)
	// 固定的、规范性的，不要轻易变更的路由
	api.RoutePreDefined(r)

	return r
}
