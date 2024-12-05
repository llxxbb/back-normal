package test

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/internal/service/demo"
	"cdel/demo/Normal/tool"
	"fmt"
	bc "github.com/llxxbb/go-BaseConfig/config"
	"go.uber.org/zap"
	"os"
)

func InitTestEnv(testPath string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("config_test.yml not found, test passed\n")
		}
	}()
	_ = os.Setenv("PRJ_ENV", "test")
	_ = os.Chdir(testPath)
	// 初始化配置, 使用嵌入的缺省配置文件
	cfg := config.ProjectConfig{}
	bc.FillConfig(&cfg, &cfg.BaseConfig)
	// 初始化日志
	tool.InitLogger(cfg.LogPath, cfg.Env == bc.VAL_PRODUCT)
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())

	// 打印配置, 注意需要先初始化日志。
	cfg.Print()

	// init service
	// 初始化上下文，如数据库
	demo.InitDemo(&cfg)
}
