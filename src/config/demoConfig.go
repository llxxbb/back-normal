package config

import (
	"github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
)

type DemoConfig struct {
	config.BaseConfig
	Mysql MysqlConfig
}

func (c *DemoConfig) AppendFieldMap(fm map[string]string) {
	c.Mysql.AppendFieldMap(fm)
}

func (c *DemoConfig) Print() {
	zap.L().Info("++++++++++++++ config info begin: ++++++++++++++")
	c.BaseConfig.Print()
	c.Mysql.Print()
	zap.L().Info("++++++++++++++ config info end: ++++++++++++++")
}
