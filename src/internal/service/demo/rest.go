package demo

import (
	"net/http"

	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/old"

	"github.com/gin-gonic/gin"
)

func V1(c *gin.Context) {
	in := old.Request{}
	c.ShouldBindJSON(&in)
	out := old.GetSuccess(in.Params)
	c.JSON(http.StatusOK, out)
}
func V2(c *gin.Context) {
	in := access.ParaIn{}
	c.ShouldBindJSON(&in)
	out := access.GetSuccessResult(in.Data)
	c.JSON(http.StatusOK, out)
}
