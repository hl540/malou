package web_server

import (
	"context"
	"github.com/google/uuid"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/hl540/malou/utils"
)

func (w *WebServer) CreateRunner(ctx context.Context, req *v1.CreateRunnerReq) (*v1.CreateRunnerResp, error) {
	runner, labels, envs := w.createRunnerReqVO2DO(req)
	err := storage.Runner.Create(ctx, runner, labels, envs)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRunnerResp{Id: runner.Id, Code: runner.Code}, nil
}

func (w *WebServer) createRunnerReqVO2DO(req *v1.CreateRunnerReq) (*storage.RunnerModel, []*storage.RunnerLabelModel, []*storage.RunnerEnvModel) {
	runner := &storage.RunnerModel{
		Code:   uuid.NewString(),
		Secret: utils.StringWithCharsetV4(20),
		Name:   req.Name,
	}
	labels := make([]*storage.RunnerLabelModel, 0, len(req.Labels))
	for _, label := range req.Labels {
		labels = append(labels, &storage.RunnerLabelModel{Value: label})
	}
	envs := make([]*storage.RunnerEnvModel, 0, len(req.Env))
	for k, v := range req.Env {
		envs = append(envs, &storage.RunnerEnvModel{Name: k, Value: v})
	}
	return runner, labels, envs
}
