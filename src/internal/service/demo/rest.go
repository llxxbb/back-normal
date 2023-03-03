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
	"github.com/go-resty/resty/v2"
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
	rows, err := config.CTX.TmpDao.SelectByName(c.Request.Context(), in.Data.(string))
	if err != nil {
		zap.S().Warn(err)
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, rows)
}
func DbSelectCached(c *gin.Context) {
	in := access.ParaIn{}
	c.ShouldBindJSON(&in)
	rows, err := config.CTX.TmpCache.SelectByName(c.Request.Context(), in.Data.(string))
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
	resp, err := getTime(config.CTX.GatewayClient)
	if err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}

	// no use here, just only show the usage of the json.Unmarshal
	raw := resp.Body()
	rtn := old.ServiceResult{}
	json.Unmarshal(raw, &rtn)
	zap.S().Debug(rtn)

	// return
	c.Data(http.StatusOK, "application/json", raw)
}
func getTime(client *resty.Client) (*resty.Response, error) {
	para := old.Request{}
	return client.R().
		SetBody(para).
		Post("/cdel@+/server/time")
}

func App2(c *gin.Context) {

	request := config.CTX.App2Client.R()
	// important! use context to chain two apps
	request.SetContext(c.Request.Context())
	rtn, _ := request.Get("/isAlive")
	c.String(http.StatusOK, rtn.String())
}
