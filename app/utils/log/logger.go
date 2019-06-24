package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func NewLogger(lp string, lv string, isDebug bool) {

	hook := lumberjack.Logger{
		Filename:   lp,    // 日志文件路径
		MaxSize:    1024,  // megabytes
		MaxBackups: 30,    // 最多保留300个备份
		MaxAge:     7,     // days
		Compress:   false, // 是否压缩 disabled by default
	}
	w := zapcore.AddSync(&hook)
	var level zapcore.Level
	switch lv {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	Logger = zap.New(core)
	Logger.Info("DefaultLogger init success")
}
