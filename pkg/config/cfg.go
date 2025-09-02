package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/codecodesam/mg/pkg/util"
)

const (
	JSON = "json"
	YAML = "yaml"
)

type MetaData struct {
	Format  string
	Content string
}

type Loader interface {
	Load() (MetaData, error)
}

type Manager struct {
	// 元数据
	container []MetaData
	// 解析结果
	parseResult map[string]any
}

var cm Manager

var once sync.Once

func GetStringValue(key string) (string, error) {
	value := cm.parseResult[key]
	if value == nil {
		return "", nil
	}
	strValue, ok := value.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("can not cast to type[%t]->[string]", value))
	}
	return strValue, nil
}

func GetIntValue(key string) (int, error) {
	value := cm.parseResult[key]
	if value == nil {
		return 0, errors.New(fmt.Sprintf("the key is not found,%s", key))
	}
	intValue, ok := value.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("the key is not int,%s", key))
	}
	return intValue, nil
}

func GetIntValueWithDefaultValue(key string, df int) int {
	value, err := GetIntValue(key)
	if err != nil {
		return df
	}
	return value
}

func NewConfigManager() {
	once.Do(func() {
		var md []MetaData
		{
			var fl Loader = FileConfigLoader{}
			fileMetaData, fileErr := fl.Load()
			if fileErr != nil {
				panic(fileErr)
			}
			md = append(md, fileMetaData)
		}
		cm.container = md
		var pr = make(map[string]any)
		for _, meta := range md {
			if meta.Format == JSON {
				m := map[string]any{}
				unmarshalErr := json.Unmarshal([]byte(meta.Content), &m)
				if unmarshalErr != nil {
					panic(unmarshalErr)
				}
				util.CopyMap(pr, m)
			} else if meta.Format == YAML {
				// TODO
			}
		}
		cm.parseResult = pr
	})

}
