package config

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/codecodesam/mg/pkg/logger"
)

type FileConfigLoader struct {
}

var (
	metaData MetaData
	err      error
)

func (loader FileConfigLoader) Load() (MetaData, error) {
	once := sync.Once{}
	// 确保只执行一次
	once.Do(func() {
		// 从环境变量读取配置文件的路径
		fp := os.Getenv("APP_CONFIG_PATH")
		if fp == "" {
			logger.Log.Error("读取不到APP_CONFIG_PATH")
			err = errors.New("读取不到APP_CONFIG_PATH")
			return
		}
		// 基于io读取文件内容最后返回配置元数据
		fn := filepath.Base(fp)
		// 按句号切开
		fs := strings.Split(fn, ".")
		// 获取文件格式
		format := fs[len(fs)-1]
		// 读取文件内容
		file, rErr := os.ReadFile(fp)
		if rErr != nil {
			logger.Log.Error("读取文件失败")
			err = errors.New("读取文件失败")
			return
		}
		// 获取文件内容
		ct := string(file)
		metaData = MetaData{format, ct}
	})
	return metaData, err
}
