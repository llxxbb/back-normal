package tool

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestClient(timeout int) *resty.Client {
	rtn := resty.New()
	rtn.SetTimeout(time.Duration(timeout) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	rtn.SetBaseURL("http://10.20.14.174/op_v2")
	return rtn
}
