package tool

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestClient(timeout int) *resty.Client {
	rtn := resty.New()
	rtn.SetTimeout(time.Duration(timeout) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	// 仅限于本 DEMO 使用
	rtn.SetHeader("HOST", "gateway.cdeledu.com")
	rtn.SetBaseURL("http://10.20.14.174/op_v2")
	// 常规方法，应放到配置中去
	// rtn.SetBaseURL("http://gateway.cdeledu.com/op_v2")
	return rtn
}
