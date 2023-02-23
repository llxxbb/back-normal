package config

import (
	"cdel/demo/Normal/tool"
	bc "github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
)

type DemoConfig struct {
	bc.BaseConfig
	Mysql       MysqlConfig
	RestTimeout int
	PinPoint    tool.PinPointConfig
}

func (c *DemoConfig) AppendFieldMap(fm map[string]string) {
	c.Mysql.AppendFieldMap(fm)
	c.PinPoint.AppendFieldMap(fm)
	fm["rest.timeOut"] = "RestTimeout"
}

func (c *DemoConfig) Print() {
	zap.L().Info("++++++++++++++ config info begin: ++++++++++++++")
	c.BaseConfig.Print()
	c.Mysql.Print()
	c.PinPoint.Print()
	zap.L().Info("------------ other setting ------------")
	zap.L().Info("-- ", zap.Int("rest.timeout", c.RestTimeout))
	zap.L().Info("++++++++++++++ config info end: ++++++++++++++")
}
