package tool

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	pphttp "github.com/pinpoint-apm/pinpoint-go-agent/plugin/http"
	"go.uber.org/zap"
)

type RpcConfig struct {
	Timeout int
	BaseUrl string
	Name    string // must equal `ProjectConfig`'s property name and so in the yaml file
}

func (c *RpcConfig) AppendFieldMap(fMap map[string]string) {
	fMap[c.Name+".timeOut"] = c.Name + ".Timeout"
	fMap[c.Name+".baseUrl"] = c.Name + ".BaseUrl"
}

func (c *RpcConfig) Print() {
	zap.L().Info(fmt.Sprintf("------------ remote process call for: %s  ------------", c.Name))
	zap.L().Info("-- ", zap.Int("timeout", c.Timeout))
	zap.L().Info("-- ", zap.String("baseUrl", c.BaseUrl))
}

func (c *RpcConfig) NewClient() *resty.Client {
	client := pphttp.WrapClient(nil) // pinpoint
	return ClientNoPP(c.Timeout, c.BaseUrl, client)
}

// ClientNoPP Compared with upper: only no pinPoint. Can be used for testing
func ClientNoPP(timeOut int, baseUrl string, client *http.Client) *resty.Client {
	rtn := resty.NewWithClient(client)
	rtn.SetTimeout(time.Duration(timeOut) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	rtn.SetBaseURL(baseUrl)
	return rtn
}
