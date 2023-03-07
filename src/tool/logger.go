package tool

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LogTmFmtWithMS = "2006-01-02 15:04:05.000"
	maxBackups     = 3
	maxAge         = 10 // 保留10天
)

func InitLogger(path string, prod bool) {

	// encoderFileCfg and debug
	encoderFileCfg := newConfig()
	encoderConsoleCfg := encoderFileCfg
	encoderConsoleCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var debugFileCore zapcore.Core
	if prod {
		debugFileCore = zapcore.NewNopCore() // 不进行输出
	} else {
		debugFileCore = newFileCore(encoderFileCfg, path+".DEBUG.log", maxBackups, maxAge, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.DebugLevel
		}))
	}

	core := zapcore.NewTee(
		// 控制台输出，常规输出
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConsoleCfg),
			// zapcore.NewJSONEncoder(encoderConsoleCfg),
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl < zapcore.ErrorLevel
			}),
		),
		// 控制台输出，错误输出
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConsoleCfg),
			zapcore.Lock(os.Stderr),
			zapcore.ErrorLevel,
		),

		// 文件输出
		newFileCore(encoderFileCfg, path+".ERROR.log", maxBackups, maxAge, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.ErrorLevel
		})),
		newFileCore(encoderFileCfg, path+".WARN.log", maxBackups, maxAge, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.WarnLevel
		})),
		newFileCore(encoderFileCfg, path+".INFO.log", maxBackups, maxAge, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel
		})),
		debugFileCore,
	)
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func newConfig() zapcore.EncoderConfig {

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(LogTmFmtWithMS))
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(caller.TrimmedPath() + ">")
	}

	// 定义输出的字段
	return zapcore.EncoderConfig{
		TimeKey:          "ts",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      zapcore.OmitKey,
		MessageKey:       "msg",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       customTimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     customCallerEncoder,
		ConsoleSeparator: " ",
	}
}

func newFileCore(encoderCfg zapcore.EncoderConfig, path string, maxBackups int, maxAge int, level zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		// zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(&lumberjack.Logger{Filename: path, MaxBackups: maxBackups, MaxAge: maxAge}),
		level,
	)
}
