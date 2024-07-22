package runner_server

import (
	"context"
	v1 "github.com/hl540/malou/proto/v1"
)

func (s *RunnerServer) RegisterRunner(ctx context.Context, req *v1.RegisterRunnerReq) (*v1.RegisterRunnerResp, error) {

	return &v1.RegisterRunnerResp{Jwt: "xxxxxxxxx"}, nil
}
