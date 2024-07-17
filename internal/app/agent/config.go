package agent

import (
	"github.com/hl540/malou/utils"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Token                 string `yaml:"Token"`                 // 认证token
	ServerAddress         string `yaml:"ServerAddress"`         // 服务地址（带端口号）
	HeartbeatFrequency    int64  `yaml:"HeartbeatFrequency"`    // 心跳频率（秒）
	PullPipelineFrequency int64  `yaml:"PullPipelineFrequency"` // 拉取流水线任务频率（秒）
	WorkerPoolSize        int    `yaml:"WorkerPoolSize"`        // 工作池大小，最大同时执行pipeline的数量
}

func LoadConfig() (*Config, error) {
	// 获取配置文件名称
	configName := utils.GetEnvDefault(ConfigNameEnvKey, "./config.yaml")

	// 加载配置文件
	configFile, err := os.Open(configName)
	if err != nil {
		return nil, err
	}

	// 解析配置文件
	var config Config
	if err = yaml.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, err
	}

	// 最终生效配置为环境变量
	config.Token = utils.GetEnvDefault(TokenEnvKey, config.Token)
	config.ServerAddress = utils.GetEnvDefault(ServerAddressEnvKey, config.ServerAddress)
	config.HeartbeatFrequency = utils.GetEnvDefault(HeartbeatFrequencyEnvKey, config.HeartbeatFrequency)
	config.PullPipelineFrequency = utils.GetEnvDefault(PullPipelineFrequencyEnvKey, config.PullPipelineFrequency)
	config.WorkerPoolSize = utils.GetEnvDefault(WorkerPoolSizeEnvKey, config.WorkerPoolSize)

	return &config, nil
}
