package runner_server

import (
	"context"
	"encoding/json"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *RunnerServer) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	runner, err := storage.Runner.GetByCode(ctx, req.Code)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err := s.saveRunnerHealth(ctx, req, runner); err != nil {
		logrus.WithContext(ctx).Warningf("save runner health failed: %v", err.Error())
	}
	return &v1.HeartbeatResp{}, nil
}

func (s *RunnerServer) saveRunnerHealth(ctx context.Context, req *v1.HeartbeatReq, runner *storage.RunnerModel) error {
	workerStatus, _ := json.Marshal(req.WorkerStatus)
	return storage.RunnerHealth.Insert(ctx, &storage.RunnerHealthModel{
		RunnerID:          runner.ID,
		CpuPercent:        req.CpuPercent,
		MemoryTotal:       req.MemoryInfo.Total,
		MemoryUsed:        req.MemoryInfo.Used,
		MemoryFree:        req.MemoryInfo.Free,
		MemoryUsedPercent: req.MemoryInfo.UsedPercent,
		DiskTotal:         req.DiskInfo.Total,
		DiskUsed:          req.DiskInfo.Used,
		DiskFree:          req.DiskInfo.Free,
		DiskUsedPercent:   req.DiskInfo.UsedPercent,
		WorkerStatus:      string(workerStatus),
		CreatedAt:         req.Timestamp,
	})
}
