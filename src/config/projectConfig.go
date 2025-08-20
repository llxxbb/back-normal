package config

import (
	"back/demo/tool"

	bc "github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
)

// PrjCfg 全局配置
var PrjCfg *ProjectConfig

type ProjectConfig struct {
	bc.BaseConfig
	Mysql    MysqlConfig
	GateWay  tool.RpcConfig
	App2     tool.RpcConfig
	PinPoint tool.PinPointConfig
	Redis    tool.RedisConfig
	Kafka    tool.KafkaConfig
}

func New() *ProjectConfig {
	PrjCfg = &ProjectConfig{}
	PrjCfg.GateWay.Name = "GateWay"
	PrjCfg.App2.Name = "App2"
	return PrjCfg
}

func (c *ProjectConfig) AppendFieldMap(fm map[string]string) {
	c.Mysql.AppendFieldMap(fm)
	c.GateWay.AppendFieldMap(fm)
	c.App2.AppendFieldMap(fm)
	c.Redis.AppendFieldMap(fm)
	c.PinPoint.AppendFieldMap(fm)
	c.Kafka.AppendFieldMap(fm)
}

func (c *ProjectConfig) Print() {
	zap.L().Info("++++++++++++++ config info begin: ++++++++++++++")
	c.BaseConfig.Print()
	c.Mysql.Print()
	c.GateWay.Print()
	c.App2.Print()
	c.Redis.Print()
	c.PinPoint.Print()
	c.Kafka.Print()
	zap.L().Info("++++++++++++++ config info end: ++++++++++++++")
}
