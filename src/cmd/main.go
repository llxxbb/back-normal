package main

import (
	"cdel/demo/Normal/api"
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/tool"
	_ "embed"
	bc "github.com/llxxbb/go-BaseConfig/config"

	"go.uber.org/zap"
)

//go:embed config_default.yml
var configDefault []byte

func main() {
	// 初始化配置, 使用嵌入的缺省配置文件
	bc.FDefault = configDefault
	cfg := config.DemoConfig{}
	bc.FillConfig(&cfg, &cfg.BaseConfig)
	cfg.Redis.Ament(&cfg.BaseConfig) // 修订 Redis Prefix
	// 初始化日志
	tool.InitLogger(cfg.LogPath, cfg.Env == bc.VAL_PRODUCT)
	defer zap.L().Sync()
	// 打印配置, 注意需要先初始化日志。
	cfg.Print()
	// 初始化上下文，如数据库
	config.CTX.Init(&cfg)

	r := api.SetupRouter()

	// start web
	r.Run(":" + cfg.Port)
}
