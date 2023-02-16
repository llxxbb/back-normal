package demo

import (
	"cdel/demo/Normal/config"
	"encoding/json"
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
		return
	}
	c.JSON(http.StatusOK, rows)
}

func DBTimeout(c *gin.Context) {
	zap.L().Info("begin")
	err := config.CTX.TmpDao.Delay()
	if err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}
	c.String(http.StatusGatewayTimeout, "O! no, it should be timeout")
}

func RemoteCall(c *gin.Context) {
	para := access.ParaIn{}
	resp, err := config.CTX.Client.R().
		SetHeader("HOST", "gateway.cdeledu.com").
		SetBody(para).
		Post("http://10.20.14.174/op_v2/cdel@+/server/time")
	if err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}
	raw := resp.Body()
	rtn := old.ServiceResult{}
	json.Unmarshal(raw, &rtn)
	zap.S().Debug(rtn)
	c.Data(http.StatusOK, "application/json", raw)
}
