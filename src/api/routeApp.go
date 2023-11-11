package api

import (
	"cdel/demo/Normal/internal/service/demo"
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
		gTmp.POST("/byName", demo.DbSelect)
		gTmp.POST("/byNameCached", demo.DbSelectCached)
		gTmp.POST("/timeout", demo.DBTimeout)
	}
	gPinPoint := r.Group("/pp")
	{
		gPinPoint.GET("/app2", demo.App2)
	}
}
