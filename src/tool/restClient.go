package tool

import (
	"time"

	"github.com/go-resty/resty/v2"
	pphttp "github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
	"go.uber.org/zap"
)

type RpcConfig struct {
	Timeout   int
	BUGateway string
	BUApp     string
}

func (c *RpcConfig) AppendFieldMap(fMap map[string]string) {
	fMap["rpc.timeOut"] = "Rpc.Timeout"
	fMap["rpc.baseUrl.gateway"] = "Rpc.BUGateway"
	fMap["rpc.baseUrl.appTwo"] = "Rpc.BUApp"
}

func (c *RpcConfig) Print() {
	zap.L().Info("------------ remote process call ------------")
	zap.L().Info("-- ", zap.Int("timeout", c.Timeout))
	zap.L().Info("-- ", zap.String("baseUrl.gateway", c.BUGateway))
	zap.L().Info("-- ", zap.String("baseUrl.appTwo", c.BUApp))
}

func RpcClient(timeOut int, baseUrl string) *resty.Client {
	client := pphttp.WrapClient(nil) // pinpoint
	rtn := resty.NewWithClient(client)
	rtn.SetTimeout(time.Duration(timeOut) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	rtn.SetBaseURL(baseUrl)
	return rtn
}
