package server

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	HttpHost      string `yaml:"HttpHost"`
	HttpPort      int    `yaml:"HttpPort"`
	GrpcHost      string `yaml:"GrpcHost"`
	GrpcPort      int    `yaml:"GrpcPort"`
	MongoUri      string `yaml:"MongoUri"`
	MongoDatabase string `yaml:"MongoDatabase"`
}

func LoadConfig() (*Config, error) {
	// 加载配置文件
	configFile, err := os.Open("./config.yaml")
	if err != nil {
		return nil, err
	}

	// 解析配置文件
	var config Config
	if err = yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
