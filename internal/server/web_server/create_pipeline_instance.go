package web_server

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) CreatePipelineInstance(ctx context.Context, req *v1.CreatePipelineInstanceReq) (*v1.CreatePipelineInstanceResp, error) {
	pipelineInstance, envs := w.createPipelineInstanceVO2DO(req)
	// 获取pipeline
	_, err := storage.Pipeline.GetById(ctx, req.PipelineId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "pipeline not found")
	}
	err = storage.PipelineInstance.Create(ctx, pipelineInstance, envs)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if req.RunnerCode != "" {
		runner, err := storage.Runner.GetByCode(ctx, req.RunnerCode)
		if err != nil {
			return nil, status.Error(codes.NotFound, "runner not found")
		}
		// 直接到指定runner
		err = w.RunPipelineInstanceOnRunner(ctx, pipelineInstance, runner.Id)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &v1.CreatePipelineInstanceResp{
			PipelineInstanceId: pipelineInstance.Id,
			Status:             v1.PipelineInstanceType_Running,
		}, nil
	}
	// 调度启动
	return &v1.CreatePipelineInstanceResp{
		PipelineInstanceId: pipelineInstance.Id,
		Status:             v1.PipelineInstanceType(pipelineInstance.Status),
	}, nil
}

func (w *WebServer) createPipelineInstanceVO2DO(req *v1.CreatePipelineInstanceReq) (*storage.PipelineInstanceModel, []*storage.PipelineInstanceEnvModel) {
	pipelineInstance := &storage.PipelineInstanceModel{
		Id:         uuid.New().String(),
		PipelineId: req.PipelineId,
	}
	envs := make([]*storage.PipelineInstanceEnvModel, 0)
	for k, v := range req.Env {
		envs = append(envs, &storage.PipelineInstanceEnvModel{Name: k, Value: v})
	}
	return pipelineInstance, envs
}

func (w *WebServer) RunPipelineInstanceOnRunner(ctx context.Context, pipelineInstance *storage.PipelineInstanceModel, runnerId int64) error {
	pipeline, steps, err := storage.Pipeline.GetInfoById(ctx, pipelineInstance.PipelineId)
	if err != nil {
		return err
	}
	runningConfig := &v1.Pipeline{
		Kind:  "docker",
		Type:  "pipeline",
		Name:  pipeline.Name,
		Steps: make([]*v1.Step, 0),
	}
	for _, step := range steps {
		runningConfig.Steps = append(runningConfig.Steps, &v1.Step{
			Name:     step.Name,
			Image:    step.Image,
			Commands: step.Commands,
		})
	}
	runningConfigByte, err := json.Marshal(runningConfig)
	if err != nil {
		return err
	}
	return storage.PipelineInstance.Start(ctx, pipelineInstance.Id, string(runningConfigByte), runnerId)
}
