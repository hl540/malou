package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) CreatePipeline(ctx context.Context, req *v1.CreatePipelineReq) (*v1.CreatePipelineResp, error) {
	pipeline, steps := w.createPipelineReq2DO(req)
	err := storage.Pipeline.Create(ctx, pipeline, steps)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.CreatePipelineResp{PipelineId: pipeline.ID}, nil
}

func (w *WebServer) createPipelineReq2DO(in *v1.CreatePipelineReq) (*storage.PipelineModel, []*storage.PipelineStepModel) {
	out1 := &storage.PipelineModel{
		Name: in.Name,
	}
	out2 := make([]*storage.PipelineStepModel, 0)
	for _, step := range in.Steps {
		out2 = append(out2, &storage.PipelineStepModel{
			Name:     step.Name,
			Image:    step.Image,
			Commands: step.Commands,
		})
	}
	return out1, out2
}
