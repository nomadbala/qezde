package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	file, _ := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(file),
		zapcore.AddSync(os.Stdout),
	)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, writer, zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}
