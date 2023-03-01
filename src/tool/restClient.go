package tool

import (
	"github.com/go-resty/resty/v2"
	"github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
	"go.uber.org/zap"
	"time"
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

func (c *RpcConfig) NewGateWayClient() *resty.Client {
	client := pphttp.WrapClient(nil) // pinpoint
	rtn := resty.NewWithClient(client)
	rtn.SetTimeout(time.Duration(c.Timeout) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	// 仅限于本 DEMO 使用
	rtn.SetHeader("HOST", "gateway.cdeledu.com")
	rtn.SetBaseURL(c.BUGateway)
	return rtn
}

// 用于测试 PinPoint 间的连通性
func (c *RpcConfig) NewApp2Client() *resty.Client {
	client := pphttp.WrapClient(nil) // pinpoint
	rtn := resty.NewWithClient(client)
	rtn.SetTimeout(time.Duration(c.Timeout) * time.Millisecond)
	rtn.SetBaseURL(c.BUApp)
	return rtn
}
