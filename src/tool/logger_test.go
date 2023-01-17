package tool

import (
	"testing"

	"go.uber.org/zap"
)

type user struct {
	name string
	age  int
}

func TestLog(t *testing.T) {
	InitLogger("/web/loges/127.0.0.1-back-normal", false)
	u := user{name: "lxb", age: 100}
	// 简易的 Sugar 方式输出，可以直接输出对象
	zap.S().Info("user info by sugarLogger: ", u)
	// 高性能的日志输出
	zap.L().Debug("log debug test")
	zap.L().Warn("log warn test")
	zap.L().Error("log error test")
	zap.L().Info("--------------user info by normalLogger ----------------",
		zap.String("name:", u.name),
		zap.Int("age:", u.age),
	)
	// defer zap.L().Sync()
}
