package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapConfig(level zap.AtomicLevel) zap.Config {
	zapConf := zap.NewProductionConfig()
	zapConf.Encoding = "console"
	zapConf.Level = level
	zapConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapConf
}
