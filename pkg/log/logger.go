package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var Send *zap.Logger

func RunLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout", "./logs.log"}
	config.ErrorOutputPaths = []string{"stderr", "./error.log"}

	var err error
	Send, err = config.Build()
	if err != nil {
		log.Fatalf("Ошибка запуска логгера: %v", err)
		return
	}
	Send.Info("Логгер запущен")
	defer Send.Sync()
}
