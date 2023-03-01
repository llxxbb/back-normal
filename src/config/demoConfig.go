package config

import (
	"cdel/demo/Normal/tool"
	bc "github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
)

type DemoConfig struct {
	bc.BaseConfig
	Mysql    MysqlConfig
	Rpc      tool.RpcConfig
	PinPoint tool.PinPointConfig
}

func (c *DemoConfig) AppendFieldMap(fm map[string]string) {
	c.Mysql.AppendFieldMap(fm)
	c.Rpc.AppendFieldMap(fm)
	c.PinPoint.AppendFieldMap(fm)
}

func (c *DemoConfig) Print() {
	zap.L().Info("++++++++++++++ config info begin: ++++++++++++++")
	c.BaseConfig.Print()
	c.Mysql.Print()
	c.Rpc.Print()
	c.PinPoint.Print()
	zap.L().Info("++++++++++++++ config info end: ++++++++++++++")
}
