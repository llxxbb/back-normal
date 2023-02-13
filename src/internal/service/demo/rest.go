package demo

import (
	"cdel/demo/Normal/config"
	"net/http"

	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"gitlab.cdel.local/platform/go/platform-common/old"
	"go.uber.org/zap"

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

func DbSelect(c *gin.Context) {
	in := access.ParaIn{}
	c.ShouldBindJSON(&in)
	rows, err := config.CTX.TmpDao.SelectByName(in.Data.(string))
	if err != nil {
		zap.S().Warn(err)
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
	}
	c.JSON(http.StatusOK, rows)
}

func DBTimeout(c *gin.Context) {
	zap.L().Info("begin")
	err := config.CTX.TmpDao.Delay()
	if err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
	}
	c.String(http.StatusOK, "O! no, it should be timeout")
}
