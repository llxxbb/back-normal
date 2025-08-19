package demo

import (
	"back/demo/internal/entity"
	"context"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/llxxbb/platform-common/access"
	"github.com/llxxbb/platform-common/def"
	"github.com/llxxbb/platform-common/old"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func V1(c *gin.Context) {
	in := old.Request[any]{}
	_ = c.ShouldBindJSON(&in)
	out := old.GetSuccess(in.Params)
	c.JSON(http.StatusOK, out)
}
func V2(c *gin.Context) {
	in := access.ParaIn[any]{}
	_ = c.ShouldBindJSON(&in)
	out := access.GetSuccessResult(in.Data)
	c.JSON(http.StatusOK, out)
}

func DbSelect(c context.Context, name string) ([]entity.TmpTable, *def.CustomError) {
	rows, err := tmpDao.SelectByName(c, name)
	if err != nil {
		zap.S().Warn(err)
		return nil, def.NewCustomError(def.ET_ENV, def.ENV_C, def.ENV_M+err.Error(), nil)
	}
	return rows, nil
}
func DbSelectCached(c *gin.Context) {
	in := access.ParaIn[string]{}
	_ = c.ShouldBindJSON(&in)
	rows, err := tmpCache.SelectByName(c.Request.Context(), in.Data)
	if err != nil {
		zap.S().Warn(err)
		c.JSON(http.StatusOK, access.GetErrorResultD[[]int](def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, rows)
}

func DBTimeout(c *gin.Context) {
	zap.L().Info("begin")
	err := tmpDao.Delay()
	if err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD[string](def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}
	c.String(http.StatusGatewayTimeout, "O! no, it should be timeout")
}

func RemoteCall(c *gin.Context) {
	resp, err := getTime(gatewayClient)
	if err != nil {
		c.JSON(http.StatusOK, access.GetErrorResultD[old.ServiceResult[any]](def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
		return
	}

	// no use here, just only show the usage of the json.Unmarshal
	raw := resp.Body()
	rtn := old.ServiceResult[any]{}
	_ = json.Unmarshal(raw, &rtn)
	zap.S().Debug(rtn)

	// return
	c.Data(http.StatusOK, "application/json", raw)
}
func getTime(client *resty.Client) (*resty.Response, error) {
	para := old.Request[any]{}
	return client.R().
		SetBody(para).
		Post("/cdel@+/server/time")
}

func App2(c *gin.Context) {
	request := app2Client.R()
	// important! use context to chain two apps
	request.SetContext(c.Request.Context())
	rtn, _ := request.Get("/isAlive")
	c.String(http.StatusOK, rtn.String())
}
