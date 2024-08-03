package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/hl540/malou/proto/v1"
)

func (w *WebServer) PipelineInfo(ctx context.Context, req *v1.PipelineInfoReq) (*v1.PipelineInfoResp, error) {
	pipeline, steps, err := storage.Pipeline.GetInfoByID(ctx, req.PipelineId)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	resp := &v1.PipelineInfoResp{
		PipelineId: pipeline.ID,
		Name:       pipeline.Name,
		Steps:      make([]*v1.Step, 0),
		CreatedAt:  pipeline.CreatedAt,
		UpdatedAt:  pipeline.UpdatedAt,
	}
	for _, step := range steps {
		resp.Steps = append(resp.Steps, w.pipelineStepDO2VO(step))
	}
	return resp, nil
}

func (w *WebServer) pipelineStepDO2VO(in *storage.PipelineStepModel) *v1.Step {
	return &v1.Step{
		Name:     in.Name,
		Image:    in.Image,
		Commands: in.Commands,
	}
}
