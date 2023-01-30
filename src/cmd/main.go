package main

import (
	"cdel/demo/Normal/api"
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/tool"

	ginzap "github.com/gin-contrib/zap"

	"go.uber.org/zap"
)

func main() {
	tool.InitLogger(config.C.LogPath, config.C.Env == config.VAL_PRODUCT)
	zap.L().Info("----------- config info begin: -----------")
	zap.S().Info(config.C)
	zap.L().Info("----------- config info end: -----------")
	defer zap.L().Sync()

	r := api.SetupRouter()
	// combine with zap
	r.Use(ginzap.Ginzap(zap.L(), tool.LogTmFmtWithMS, false))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// start web
	r.Run(":" + config.C.Port)
}
