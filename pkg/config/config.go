package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Instance *Config // 全局变量

type Config struct {
	Env      string `yaml:"Env"`  // 环境：prod生产环境、env开发
	Port     string `yaml:"Port"` // 服务端口
	MySqlUrl string `yaml:"MySqlUrl"`
	// redis
	// elastic
	// kafka
}

// 初始化全局配置：从yaml文件读取并存到Config实例里面
func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		// 读取配置文件失败
		panic("读取配置文件失败")
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		// Unmarshal失败
		panic("Unmarshal配置文件失败")
	}
	return Instance
}
