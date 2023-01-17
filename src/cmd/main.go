package main

import (
	"cdel/demo/Normal/api"
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/tool"

	"go.uber.org/zap"
)

func main() {
	tool.InitLogger(config.C.LogPath, config.C.Env == config.VAL_PRODUCT)
	zap.S().Info("project config: ", config.C)
	defer zap.L().Sync()

	r := api.SetupRouter()
	r.Run(":" + config.C.Port)
}
