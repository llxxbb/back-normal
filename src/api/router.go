package api

import (
	"cdel/demo/Normal/config"
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

	// 设置中间间-----------------------------------
	// combine with zap
	//r.Use(ginzap.Ginzap(zap.L(), tool.LogTmFmtWithMS, false))		// 可以打印所有的请求日志
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// use pinpoint
	r.Use(ppgin.Middleware())

	// 路由信息 -----------------------------------------------
	// 需要自行定制的路由
	routeApp(r)
	// 固定的、规范性的，不要轻易变更的路由
	routePreDefined(r)

	return r
}
