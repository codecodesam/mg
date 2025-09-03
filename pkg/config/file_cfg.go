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
	metaData       MetaData
	err            error
	fileLoaderOnce sync.Once
)

func (loader FileConfigLoader) Load() (MetaData, error) {
	fileLoaderOnce.Do(func() {
		fp := os.Getenv(ENV_APP_CONFIG_PATH)
		if fp == "" {
			logger.Log.Error("can not find env APP_CONFIG_PATH")
			err = errors.New("can not find env APP_CONFIG_PATH")
			return
		}
		fn := filepath.Base(fp)
		fs := strings.Split(fn, ".")
		format := fs[len(fs)-1]
		file, rErr := os.ReadFile(fp)
		if rErr != nil {
			logger.Log.Error("read config file error", rErr)
			err = errors.New("read config file error")
			return
		}
		ct := string(file)
		metaData = MetaData{format, ct}
	})
	return metaData, err
}
