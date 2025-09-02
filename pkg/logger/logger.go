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
	// 配置 Lumberjack 进行日志切割
	lumberJackLogger := &lumberjack.Logger{
		// 日志文件路径，如 "./logs/app.log"
		Filename: logPath,
		// 单个文件最大尺寸（单位：MB），超过则切割
		MaxSize: 10,
		// 保留旧文件的最大个数
		MaxBackups: 5,
		// 保留旧文件的最大天数
		MaxAge: 30,
		// 是否压缩/归档旧文件
		Compress: true,
	}

	// 创建多个 Writer，例如同时输出到文件和标准输出
	fileDebugging := zapcore.AddSync(lumberJackLogger)
	consoleDebugging := zapcore.AddSync(os.Stdout)

	// 编码器配置，定义日志格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime) // ISO8601 时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig) // 使用 JSON 格式输出

	// 创建 Core，可以组合多个输出目的和级别
	core := zapcore.NewTee(
		// 所有级别的日志到文件
		zapcore.NewCore(encoder, fileDebugging, zap.DebugLevel),
		// 所有级别的日志到控制台
		zapcore.NewCore(encoder, consoleDebugging, zap.DebugLevel),
	)

	// Logger，并添加调用者信息和堆栈跟踪
	//zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	zapLogger := zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
	// 转换为 SugaredLogger，支持格式化日志
	Log = zapLogger.Sugar()
}
