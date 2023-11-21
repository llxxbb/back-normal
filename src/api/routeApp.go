package api

import (
	"cdel/demo/Normal/internal/service/demo"
	"cdel/demo/Normal/tool"
	"github.com/gin-gonic/gin"
)

func routeApp(r *gin.Engine) {
	gDemo := r.Group("/demo")
	{
		gDemo.POST("/v1", demo.V1)
		gDemo.POST("/v2", demo.V2)
		gDemo.POST("/rest", demo.RemoteCall)
	}
	gTmp := r.Group("/tmp")
	{
		// 注意输入参数仅适用于 `access.ParaIn`，DbSelect 不需要进行参数处理。推荐！
		tool.RequestResponse(gTmp, "byName", demo.DbSelect)
		// 常规处理方式，DbSelectCached 需要自行处理参数。，
		gTmp.POST("/byNameCached", demo.DbSelectCached)
		gTmp.POST("/timeout", demo.DBTimeout)
	}
	gPinPoint := r.Group("/pp")
	{
		gPinPoint.GET("/app2", demo.App2)
	}
}
