package runner_server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var i = 0

func (s *RunnerServer) PullPipeline(ctx context.Context, req *v1.PullPipelineReq) (*v1.PullPipelineResp, error) {
	if i > 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "pull_pipeline is already running")
	}
	i++
	fmt.Printf("[PullPipeline] %v\n", req)
	runner, err := storage.Runner.GetByCode(ctx, "f063f1f1-8170-4d87-9807-8bae7c991394")
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	pipelineInstance, err := storage.PipelineInstance.GetPendingByRunnerId(ctx, runner.Id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, status.Error(codes.NotFound, "pipeline not found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var runningConfig v1.Pipeline
	err = json.Unmarshal([]byte(pipelineInstance.RuntimeConfig), &runningConfig)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.PullPipelineResp{
		PipelineId: pipelineInstance.Id,
		Pipeline:   &runningConfig,
	}, nil
}
