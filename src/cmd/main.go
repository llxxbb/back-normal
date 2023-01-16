package main

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/tool"

	"go.uber.org/zap"
)

func main() {
	tool.InitLogger(config.C.LogPath, config.C.Env == config.VAL_PRODUCT)
	zap.S().Info("project config: ", config.C)
	zap.L().Debug("log debug test")
	zap.L().Warn("log warn test")
	zap.L().Error("log error test")
	zap.L().WithOptions()
	defer zap.L().Sync()
}
