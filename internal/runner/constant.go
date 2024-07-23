package runner

const (
	ConfigNameEnvKey            = "CONFIG_NAME"
	TokenEnvKey                 = "TOKEN"
	JwtFileEnvKey               = "JWT_FILE"
	ServerHostEnvKey            = "SERVER_HOST"
	ServerPortEnvKey            = "SERVER_PORT"
	HeartbeatFrequencyEnvKey    = "HEARTBEAT_FREQUENCY"
	PullPipelineFrequencyEnvKey = "PULL_PIPELINE_FREQUENCY"
	WorkerPoolSizeEnvKey        = "WORKER_POOL_SIZE"
	WorkDirEnvKey               = "WORK_DIR"
)

const ServerPortDefault = 5555
const JwtFileDefault = ".jwt"
const HeartbeatFrequencyDefault = 30
const PullPipelineFrequencyDefault = 15
const WorkerPoolSizeDefault = 2
