package runner

import (
	"github.com/hl540/malou/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Token                 string `yaml:"Token"`                 // token
	JwtFile               string `yaml:"JwtFile"`               // 注册完成后生成的jwt
	ServerHost            string `yaml:"ServerHost"`            // 服务地址
	ServerPort            int    `yaml:"ServerPort"`            // 服务端口
	HeartbeatFrequency    int64  `yaml:"HeartbeatFrequency"`    // 心跳频率（秒）
	PullPipelineFrequency int64  `yaml:"PullPipelineFrequency"` // 拉取流水线任务频率（秒）
	WorkerPoolSize        int    `yaml:"WorkerPoolSize"`        // 工作池大小，最大同时执行pipeline的数量
	WorkDir               string `yaml:"WorkDir"`               // 工作目录
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

	//加载环境变量
	if err := godotenv.Load(); err != nil {
		logrus.Warning("Didn't try and open .env by default")
	}

	// 最终生效配置为环境变量
	config.Token = utils.GetEnvDefault(TokenEnvKey, config.Token)
	config.JwtFile = utils.GetEnvDefault(JwtFileEnvKey, config.JwtFile)
	config.ServerHost = utils.GetEnvDefault(ServerHostEnvKey, config.ServerHost)
	config.ServerPort = utils.GetEnvDefault(ServerPortEnvKey, config.ServerPort)
	config.HeartbeatFrequency = utils.GetEnvDefault(HeartbeatFrequencyEnvKey, config.HeartbeatFrequency)
	config.PullPipelineFrequency = utils.GetEnvDefault(PullPipelineFrequencyEnvKey, config.PullPipelineFrequency)
	config.WorkerPoolSize = utils.GetEnvDefault(WorkerPoolSizeEnvKey, config.WorkerPoolSize)
	config.WorkDir = utils.GetEnvDefault(WorkDirEnvKey, config.WorkDir)

	// 默认值
	if config.ServerPort == 0 {
		config.ServerPort = ServerPortDefault
	}
	if config.JwtFile == "" {
		config.JwtFile = JwtFileDefault
	}
	if config.HeartbeatFrequency == 0 {
		config.HeartbeatFrequency = HeartbeatFrequencyDefault
	}
	if config.PullPipelineFrequency == 0 {
		config.PullPipelineFrequency = PullPipelineFrequencyDefault
	}

	return &config, nil
}
