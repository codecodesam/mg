package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

func InitLogger(logPath string) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	fileDebugging := zapcore.AddSync(lumberJackLogger)
	consoleDebugging := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileDebugging, zap.DebugLevel),
		zapcore.NewCore(encoder, consoleDebugging, zap.DebugLevel),
	)

	zapLogger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
	Log = zapLogger.Sugar()
}
