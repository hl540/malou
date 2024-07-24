package runner_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"github.com/sirupsen/logrus"
)

func (s *RunnerServer) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	if err := s.saveRunnerHealth(ctx, req); err != nil {
		logrus.WithContext(ctx).Warningf("save runner health failed: %v", err.Error())
	}
	return &v1.HeartbeatResp{}, nil
}

func (s *RunnerServer) saveRunnerHealth(ctx context.Context, req *v1.HeartbeatReq) error {
	return storage.AddRunnerHealth(ctx, &storage.RunnerHealth{
		RunnerID:   req.Token,
		CreatedAt:  req.Timestamp,
		CpuPercent: req.CpuPercent,
		MemoryInfo: &storage.DiskInfo{
			Total:       req.MemoryInfo.Total,
			Used:        req.MemoryInfo.Used,
			Free:        req.MemoryInfo.Free,
			UsedPercent: req.MemoryInfo.UsedPercent,
		},
		DiskInfo: &storage.DiskInfo{
			Total:       req.DiskInfo.Total,
			Used:        req.DiskInfo.Used,
			Free:        req.DiskInfo.Free,
			UsedPercent: req.DiskInfo.UsedPercent,
		},
		WorkerStatus: req.WorkerStatus,
	})
}
