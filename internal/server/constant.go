package server

const (
	ConfigNameEnvKey            = "CONFIG_NAME"
	TokenEnvKey                 = "TOKEN"
	ServerAddressEnvKey         = "SERVER_ADDRESS"
	HeartbeatFrequencyEnvKey    = "HEARTBEAT_FREQUENCY"
	PullPipelineFrequencyEnvKey = "PULL_PIPELINE_FREQUENCY"
	WorkerPoolSizeEnvKey        = "WORKER_POOL_SIZE"
	WorkDirEnvKey               = "WORK_DIR"
)

const HeartbeatFrequencyDefault = 30
const PullPipelineFrequencyDefault = 15
