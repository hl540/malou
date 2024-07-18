package runner

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/hl540/malou/internal/runner/container_runtime"
	"github.com/hl540/malou/internal/runner/worker"
	"github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"path"
	"time"
)

var Logger = logrus.New()

type Runner struct {
	Token        string
	Config       *Config
	DockerClient *client.Client
	MalouClient  v1.MalouClient
}

func NewRunner(conf *Config) (*Runner, error) {
	app := &Runner{
		Token:  conf.Token,
		Config: conf,
	}

	var err error
	// 初始化docker cli
	app.DockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("initialize docker cli failed, %s", err.Error())
	}

	// 初始化服务器grpc client
	conn, err := grpc.NewClient(conf.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("initialize server grpc client failed, %s", err.Error())
	}
	app.MalouClient = v1.NewMalouClient(conn)

	return app, nil
}

func (a *Runner) Run(ctx context.Context) {
	heartbeatTicker := time.NewTicker(time.Duration(a.Config.HeartbeatFrequency) * time.Second)
	pullPipelineTicker := time.NewTicker(time.Duration(a.Config.PullPipelineFrequency) * time.Second)
	for {
		select {
		case <-heartbeatTicker.C:
			a.Heartbeat(ctx)
		case <-pullPipelineTicker.C:
			a.PullPipeline(ctx)
		case <-ctx.Done():
			heartbeatTicker.Stop()
			pullPipelineTicker.Stop()
			return
		}
	}
}

// Heartbeat 心跳，与服务端保持联系，并拉取服务端指令
func (a *Runner) Heartbeat(ctx context.Context) {
	heartbeatResp, err := a.MalouClient.Heartbeat(ctx, &v1.HeartbeatReq{
		Token:        a.Token,
		CpuPercent:   utils.GetCpuPercent(),
		MemoryInfo:   utils.GetMemoryPercent(),
		DiskInfo:     utils.GetDiskPercent(),
		WorkerStatus: worker.Pool.Status(),
	})
	if err != nil {
		Logger.WithContext(ctx).Errorf("[Heartbeat] request failed, err: %s", err.Error())
		return
	}
	Logger.WithContext(ctx).Infof("[Heartbeat] %d %s", heartbeatResp.Timestamp, heartbeatResp.Message)
}

func (a *Runner) PullPipeline(ctx context.Context) {
	// 尝试获取work
	workID := worker.Pool.TryWorker()
	if workID == "" {
		Logger.WithContext(ctx).Info("[PullPipeline] there are no idle worker")
		return
	}

	// 拉取Pipeline
	pullPipelineResp, err := a.MalouClient.PullPipeline(ctx, &v1.PullPipelineReq{})
	if err != nil {
		// 归还worker
		worker.Pool.Release(workID)
		Logger.WithContext(ctx).Errorf("[PullPipeline] request failed, err: %s", err.Error())
		return
	}
	if pullPipelineResp.PipelineId == "" {
		// 归还worker
		worker.Pool.Release(workID)
		Logger.WithContext(ctx).Infof("[PullPipeline] No pull, %s", pullPipelineResp.Message)
		return
	}
	// 使用拉取到的PipelineID填充work
	if !worker.Pool.Worker(workID, pullPipelineResp.PipelineId) {
		Logger.WithContext(ctx).Info("[PullPipeline] worker don't exist")
		return
	}
	newCtx := context.Background()
	go func() {
		a.ExecutePipeline(newCtx, pullPipelineResp.PipelineId, pullPipelineResp.Pipeline)
		// 归还worker
		worker.Pool.Release(workID)
	}()
}

func (a *Runner) ExecutePipeline(ctx context.Context, pipelineID string, pipeline *v1.Pipeline) {
	// 创建stream，回传执行过程log
	reportStream, err := a.MalouClient.ReportPipelineLog(ctx)
	if err != nil {
		Logger.WithContext(ctx).Errorf("Failed to create report stream, %s", err.Error())
	}
	defer reportStream.CloseAndRecv()
	reportLog := NewReportLog(pipelineID, reportStream)

	// 创建执行step的环境，默认docker容器运行时
	containerRuntime, err := container_runtime.NewDockerRuntime(a.DockerClient)
	if err != nil {
		reportLog.Log("Failed to create container runtime, %s", err.Error())
		return
	}

	// 工作目录，需要挂载到容器中
	workDir := path.Join(a.Config.WorkDir, pipelineID)

	reportLog.Log("start executing pipeline [%s],[%s]", pipeline.Name, pipelineID)
	for _, step := range pipeline.Steps {
		// 创建step执行器
		stepExecutor := NewBaseStepExecutor(containerRuntime, reportLog)
		if err := stepExecutor.Execute(ctx, step, workDir); err != nil {
			reportLog.Error("execution failure: %s", err.Error())
			return
		}
	}
	reportLog.Done("complete")
}
