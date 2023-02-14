package venus

import (
	"github.com/no-mole/venus/agent/venus/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapConfig(conf *config.Config) zap.Config {
	zapConf := zap.NewProductionConfig()
	zapConf.Encoding = "console"
	zapConf.Level = conf.ZapLoggerLevel()
	zapConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapConf
}
