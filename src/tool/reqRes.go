package tool

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/llxxbb/platform-common/access"
	"github.com/llxxbb/platform-common/def"
	"github.com/llxxbb/platform-common/old"
	"go.uber.org/zap"
)

// RequestResponse 此方法因涉及到Gin框架，因此没有封装到基础工具包中。
func RequestResponse[T any, R any](rg *gin.RouterGroup, path string, fun func(c context.Context, p T) (R, *def.CustomError)) {
	rg.POST(path, func(c *gin.Context) {
		zap.L().Debug("accessed", zap.String("url", c.Request.URL.Path))
		in := access.ParaIn[T]{}
		_ = c.ShouldBindJSON(&in)
		out := access.GetResultByParaCtx(c.Request.Context(), in.Data, fun)
		c.JSON(http.StatusOK, out)
	})
}

func RequestResponseOld[T any, R any](rg *gin.RouterGroup, path string, fun func(c context.Context, p T) (R, *def.CustomError)) {
	rg.POST(path, func(c *gin.Context) {
		zap.L().Debug("accessed", zap.String("url", c.Request.URL.Path))
		in := old.Request[T]{}
		_ = c.ShouldBindJSON(&in)
		out := old.GetResultByParaCtx(c.Request.Context(), in.Params, fun)
		c.JSON(http.StatusOK, out)
	})
}
