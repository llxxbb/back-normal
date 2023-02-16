package config

import (
	"github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
)

type DemoConfig struct {
	config.BaseConfig
	Mysql       MysqlConfig
	RestTimeout int
}

func (c *DemoConfig) AppendFieldMap(fm map[string]string) {
	c.Mysql.AppendFieldMap(fm)
	fm["rest.timeOut"] = "RestTimeout"
}

func (c *DemoConfig) Print() {
	zap.L().Info("++++++++++++++ config info begin: ++++++++++++++")
	c.BaseConfig.Print()
	c.Mysql.Print()
	zap.L().Info("------------ other setting ------------")
	zap.L().Info("-- ", zap.Int("rest.timeout", c.RestTimeout))
	zap.L().Info("++++++++++++++ config info end: ++++++++++++++")
}
