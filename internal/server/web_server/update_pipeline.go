package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) UpdatePipeline(ctx context.Context, req *v1.UpdatePipelineReq) (*v1.UpdatePipelineResp, error) {
	newPipeline, newSteps := w.updatePipelineReq2DO(req)
	pipeline, err := storage.Pipeline.GetByID(ctx, newPipeline.ID)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	err = storage.Pipeline.Update(ctx, newPipeline, newSteps)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &v1.UpdatePipelineResp{PipelineId: pipeline.ID}, nil
}

func (w *WebServer) updatePipelineReq2DO(in *v1.UpdatePipelineReq) (*storage.PipelineModel, []*storage.PipelineStepModel) {
	out1 := &storage.PipelineModel{
		ID:   in.PipelineId,
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
