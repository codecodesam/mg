package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
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
		return "", errors.New(fmt.Sprintf("无法转换值类型[%t]->[string]", value))
	}
	return strValue, nil
}

func NewConfigManager() {
	once.Do(func() {
		var md []MetaData
		// 加载文件配置
		{
			var fl Loader = FileConfigLoader{}
			// 加载逻辑
			fileMetaData, fileErr := fl.Load()
			if fileErr != nil {
				panic(fileErr)
			}
			md = append(md, fileMetaData)
		}
		// 设置属性
		cm.container = md
		// 解析属性
		var pr map[string]any
		//
		for _, meta := range md {
			if meta.Format == JSON {
				// 解析json
				m := map[string]any{}
				// json序列化
				unmarshalErr := json.Unmarshal([]byte(meta.Content), &m)
				if unmarshalErr != nil {
					panic(unmarshalErr)
				}
				// TODO 追加map

			} else if meta.Format == YAML {
				// 解析yaml
			}
		}
		cm.parseResult = pr
	})

}
