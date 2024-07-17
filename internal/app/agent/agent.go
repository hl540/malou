package agent

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/hl540/malou/internal/container_runtime"
	"github.com/hl540/malou/internal/worker"
	"github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Agent struct {
	Token        string
	Config       *Config
	DockerClient *client.Client
	MalouClient  v1.MalouServerClient
}

func NewAgent(conf *Config) (*Agent, error) {
	agent := &Agent{
		Token:  conf.Token,
		Config: conf,
	}

	var err error
	// 初始化docker cli
	agent.DockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("initialize docker cli failed, %s", err.Error())
	}

	// 初始化服务器grpc client
	conn, err := grpc.NewClient(conf.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("initialize server grpc client failed, %s", err.Error())
	}
	agent.MalouClient = v1.NewMalouServerClient(conn)

	return agent, nil
}

func (a *Agent) Run(ctx context.Context) {
	heartbeatTicker := time.NewTicker(time.Duration(a.Config.HeartbeatFrequency) * time.Second)
	pullPipelineTicker := time.NewTicker(time.Duration(a.Config.PullPipelineFrequency) * time.Second)
	debugTicker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-heartbeatTicker.C:
			a.Heartbeat(ctx)
		case <-pullPipelineTicker.C:
			a.PullPipeline(ctx)
		case <-debugTicker.C:
			fmt.Printf("-----------------------------work status %v\n", worker.Worker.Status())
		case <-ctx.Done():
			heartbeatTicker.Stop()
			pullPipelineTicker.Stop()
			return
		}
	}
}

// Heartbeat 心跳，与服务端保持联系，并拉取服务端指令
func (a *Agent) Heartbeat(ctx context.Context) {
	heartbeatResp, err := a.MalouClient.Heartbeat(ctx, &v1.HeartbeatReq{
		AgentToken:   a.Token,
		CpuPercent:   utils.GetCpuPercent(),
		MemoryInfo:   utils.GetMemoryPercent(),
		DiskInfo:     utils.GetDiskPercent(),
		WorkerStatus: worker.Worker.Status(),
	})
	if err != nil {
		Logger.WithContext(ctx).Errorf("[Heartbeat] request failed, err: %s", err.Error())
		return
	}
	Logger.WithContext(ctx).Infof("[Heartbeat] %d %s", heartbeatResp.Timestamp, heartbeatResp.Message)
}

func (a *Agent) PullPipeline(ctx context.Context) {
	// 尝试获取work
	workID := worker.Worker.TryWorker()
	if workID == "" {
		Logger.WithContext(ctx).Info("[PullPipeline] there are no idle worker")
		return
	}
	defer worker.Worker.Release(workID)

	// 拉取Pipeline
	pullPipelineResp, err := a.MalouClient.PullPipeline(ctx, &v1.PullPipelineReq{})
	if err != nil {
		Logger.WithContext(ctx).Errorf("[PullPipeline] request failed, err: %s", err.Error())
		return
	}
	if pullPipelineResp.PipelineId == "" {
		Logger.WithContext(ctx).Infof("[PullPipeline] No pull, %s", pullPipelineResp.Message)
		return
	}
	// 使用拉取到的PipelineID填充work
	if !worker.Worker.Worker(workID, pullPipelineResp.PipelineId) {
		Logger.WithContext(ctx).Info("[PullPipeline] worker don't exist")
		return
	}
	newCtx := context.Background()
	go a.ExecutePipeline(newCtx, pullPipelineResp.PipelineId, pullPipelineResp.Pipeline)
}

func (a *Agent) ExecutePipeline(ctx context.Context, pipelineID string, pipeline *v1.Pipeline) {
	// 工作目录，需要挂载到容器中
	workDir := "C:/Users/67440/go/src/malou/work/" + pipelineID

	// 创建上报log流
	reportStream, err := a.MalouClient.ReportPipelineLog(ctx)
	if err != nil {
		// TODO
		return
	}
	defer reportStream.CloseAndRecv()
	reportLog := NewReportLog(pipelineID, reportStream)

	containerRuntime, err := container_runtime.NewDockerRuntime(a.DockerClient)
	if err != nil {
		return
	}
	for _, step := range pipeline.Steps {
		reportLog.Log("Start execution %s", step.Name)
		stepReportLog := reportLog.WithStep(step.Name)
		// 创建step执行器
		stepExecutor := NewBaseStepExecutor(containerRuntime, stepReportLog)
		if err := stepExecutor.Execute(ctx, step, workDir); err != nil {
			reportLog.Error("execution failure: ", err.Error())
			return
		}
	}
	reportLog.Log("complete")
}
