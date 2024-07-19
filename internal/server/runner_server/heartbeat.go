package runner_server

import (
	"context"
	v1 "github.com/hl540/malou/proto/v1"
	"time"
)

func (s *RunnerServer) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	//Logger.WithContext(ctx).Infof("[Heartbeat] %v", req)
	return &v1.HeartbeatResp{Timestamp: time.Now().Unix(), Message: "Received"}, nil
}
