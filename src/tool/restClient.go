package tool

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestClient(timeout int) *resty.Client {
	rtn := resty.New()
	rtn.SetTimeout(time.Duration(timeout) * time.Millisecond)
	rtn.SetHeader("Content-Type", "application/json")
	return rtn
}
