package config

import (
	"github.com/BurntSushi/toml"
)

var conf map[string]map[string]interface{}

// ParseConfig 解析配置文件
func ParseConfig(filePath string) error {
	_, err := toml.DecodeFile(filePath, &conf)
	return err
}

// GetConfig 获取配置信息
func GetConfig(configName string) map[string]interface{} {
	return conf[configName]
}
