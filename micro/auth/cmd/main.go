package main

import (
	"os"
	"time"

	"github.com/codecodesam/mg/pkg/config"
	"github.com/codecodesam/mg/pkg/logger"
)

// TODO 文件没生成
func main() {
	// 日志初始化
	initLogger()
	// 初始化配置
	initConfig()

	logger.Log.Info("test...")

	closeLogger()

	time.Sleep(10 * time.Second)
}

func initConfig() {
	config.NewConfigManager()
}

func initLogger() {
	path := os.Getenv("LOGGER_PATH")
	if path == "" {
		panic("环境变量中找不到日志的路径")
	}
	logger.InitLogger(path)
}

func closeLogger() {
	logger.Log.Sync()
}
