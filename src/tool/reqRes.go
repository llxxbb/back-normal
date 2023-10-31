package tool

import (
	"github.com/gin-gonic/gin"
	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"go.uber.org/zap"
	"net/http"
)

func RequestResponse[T any, R any](rg *gin.RouterGroup, path string, fun func(p T) (R, *def.CustomError)) {
	rg.POST(path, func(c *gin.Context) {
		zap.L().Debug("accessed", zap.String("url", c.Request.URL.Path))
		in := access.ParaIn[T]{}
		_ = c.ShouldBindJSON(&in)
		out := access.GetResultWithParam(in.Data, fun)
		c.JSON(http.StatusOK, out)
	})
}
