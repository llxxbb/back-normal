package config

import (
	"cdel/demo/Normal/tool"

	bc "github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
)

// PrjCfg 全局配置
var PrjCfg *ProjectConfig

type ProjectConfig struct {
	bc.BaseConfig
	Mysql    MysqlConfig
	Rpc      tool.RpcConfig
	PinPoint tool.PinPointConfig
	Redis    tool.RedisConfig
}

func New() *ProjectConfig {
	PrjCfg = &ProjectConfig{}
	return PrjCfg
}

func (c *ProjectConfig) AppendFieldMap(fm map[string]string) {
	c.Mysql.AppendFieldMap(fm)
	c.Rpc.AppendFieldMap(fm)
	c.Redis.AppendFieldMap(fm)
	c.PinPoint.AppendFieldMap(fm)
}

func (c *ProjectConfig) Print() {
	zap.L().Info("++++++++++++++ config info begin: ++++++++++++++")
	c.BaseConfig.Print()
	c.Mysql.Print()
	c.Rpc.Print()
	c.Redis.Print()
	c.PinPoint.Print()
	zap.L().Info("++++++++++++++ config info end: ++++++++++++++")
}
