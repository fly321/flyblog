package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type DatabaseConfig struct {
	Drive    string `yaml:"drive"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Config struct {
	Databases []DatabaseConfig `yaml:"databases"`
}

var DbConfig Config

// init 读取配置文件
func init() {
	// 读取配置文件
	// 读取配置文件
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	// 解析配置文件
	if err := yaml.Unmarshal(data, &DbConfig); err != nil {
		log.Fatalf("无法解析配置文件: %v", err)
	}
}

func (receiver Config) ListDb() Config {
	return receiver
}
