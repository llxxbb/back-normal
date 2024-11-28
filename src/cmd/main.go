package main

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/internal/service/demo"
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
	cfg := config.New()
	bc.FillConfig(cfg, &cfg.BaseConfig)
	cfg.Redis.Amend(&cfg.BaseConfig) // 修订 Redis Prefix
	// 初始化日志
	tool.InitLogger(cfg.LogPath, cfg.Env == bc.VAL_PRODUCT)
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())

	// 打印配置, 注意需要先初始化日志。
	cfg.Print()

	// init pinpoint
	cfg.PinPoint.InitPinPoint(cfg.ProjectName, cfg.Host)
	// 初始化上下文，如数据库
	demo.InitDemo(cfg)

	r := SetupRouter(cfg)

	// start web
	_ = r.Run(":" + cfg.Port)
}
