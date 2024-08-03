package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) PipelineList(ctx context.Context, req *v1.PipelineListReq) (*v1.PipelineListResp, error) {
	list, count, err := storage.Pipeline.SearchList(ctx, &storage.PipelineSearchListParam{
		Name: req.Name,
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	data := make([]*v1.PipelineListItem, 0, len(list))
	for _, pipeline := range list {
		data = append(data, w.pipelineListItemDO2VO(pipeline))
	}
	return &v1.PipelineListResp{
		Total: count,
		Data:  data,
	}, nil
}

func (w *WebServer) pipelineListItemDO2VO(in *storage.PipelineModel) *v1.PipelineListItem {
	return &v1.PipelineListItem{
		PipelineId: in.ID,
		Name:       in.Name,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}
}
