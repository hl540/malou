package web_server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/hl540/malou/internal/server/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/hl540/malou/proto/v1"
)

func (w *WebServer) RunnerInfo(ctx context.Context, req *v1.RunnerInfoReq) (*v1.RunnerInfoResp, error) {
	runner, err := storage.Runner.GetByID(ctx, req.RunnerId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	labels, err := storage.Runner.GetLabelByRunnerID(ctx, runner.Code)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	envs, err := storage.Runner.GetEnvByRunnerID(ctx, runner.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	health, err := storage.RunnerHealth.GetLatestByRunnerID(ctx, runner.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := &v1.RunnerInfoResp{
		Name:      runner.Name,
		Code:      runner.Code,
		Status:    v1.RunnerStatusType(runner.Status),
		Labels:    w.runnerLabelDO2VO(labels),
		Env:       w.runnerEnvDO2VO(envs),
		CreatedAt: runner.CreatedAt,
		UpdatedAt: runner.UpdatedAt,
		Health:    w.runnerHealthDO2VO(health),
	}
	return resp, nil
}

func (w *WebServer) runnerLabelDO2VO(in []*storage.RunnerLabelModel) []string {
	if len(in) == 0 {
		return make([]string, 0)
	}
	out := make([]string, 0, len(in))
	for _, label := range in {
		out = append(out, label.Value)
	}
	return out
}

func (w *WebServer) runnerEnvDO2VO(in []*storage.RunnerEnvModel) map[string]string {
	if len(in) == 0 {
		return make(map[string]string)
	}
	out := make(map[string]string, len(in))
	for _, env := range in {
		out[env.Name] = env.Value
	}
	return out
}

func (w *WebServer) runnerHealthDO2VO(in *storage.RunnerHealthModel) *v1.RunnerInfoHealth {
	if in == nil {
		return nil
	}
	var workerStatus map[string]string
	_ = json.Unmarshal([]byte(in.WorkerStatus), &workerStatus)
	return &v1.RunnerInfoHealth{
		CpuPercent: in.CpuPercent,
		MemoryInfo: &v1.MemoryInfo{
			Total:       in.MemoryTotal,
			Used:        in.MemoryUsed,
			Free:        in.MemoryFree,
			UsedPercent: in.MemoryUsedPercent,
		},
		DiskInfo: &v1.DiskInfo{
			Total:       in.DiskTotal,
			Used:        in.DiskUsed,
			Free:        in.DiskFree,
			UsedPercent: in.DiskUsedPercent,
		},
		WorkerStatus: workerStatus,
	}
}
