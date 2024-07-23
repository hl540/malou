package runner_server

import (
	"context"
	v1 "github.com/hl540/malou/proto/v1"
)

func (s *RunnerServer) Heartbeat(ctx context.Context, req *v1.HeartbeatReq) (*v1.HeartbeatResp, error) {
	//Logger.WithContext(ctx).Infof("[Heartbeat] %v", req)
	return &v1.HeartbeatResp{Code: 0, Message: "Received"}, nil
}
