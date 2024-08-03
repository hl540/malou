package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) UpdateRunner(ctx context.Context, req *v1.UpdateRunnerReq) (*v1.UpdateRunnerResp, error) {
	newRunner, newLabels, newEnvs := w.updateRunnerReqVO2DO(req)
	runner, err := storage.Runner.GetByID(ctx, newRunner.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	newRunner.Id = runner.Id
	newRunner.Code = runner.Code
	err = storage.Runner.Update(ctx, newRunner, newLabels, newEnvs)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.UpdateRunnerResp{Id: runner.Id, Code: runner.Code}, nil
}

func (w *WebServer) updateRunnerReqVO2DO(req *v1.UpdateRunnerReq) (*storage.RunnerModel, []*storage.RunnerLabelModel, []*storage.RunnerEnvModel) {
	runner := &storage.RunnerModel{
		Id:   req.Id,
		Name: req.Name,
	}
	labels := make([]*storage.RunnerLabelModel, 0, len(req.Labels))
	for _, label := range req.Labels {
		labels = append(labels, &storage.RunnerLabelModel{RunnerId: req.Id, Value: label})
	}
	envs := make([]*storage.RunnerEnvModel, 0, len(req.Env))
	for k, v := range req.Env {
		envs = append(envs, &storage.RunnerEnvModel{RunnerId: req.Id, Name: k, Value: v})
	}
	return runner, labels, envs
}
