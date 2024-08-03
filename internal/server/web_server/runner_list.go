package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) RunnerList(ctx context.Context, req *v1.RunnerListReq) (*v1.RunnerListResp, error) {
	// 获取runner列表
	list, count, err := storage.Runner.SearchList(ctx, &storage.RunnerSearchListParam{
		Code:   req.Code,
		Name:   req.Name,
		Labels: req.Labels,
		Page:   req.Page,
		Size:   req.Size,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	data := make([]*v1.RunnerListItem, 0, len(list))
	for _, runner := range list {
		labels, err := storage.Runner.GetLabelsByRunnerId(ctx, runner.Id)
		if err != nil {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		data = append(data, w.runnerListItemDO2VO(runner, labels))
	}
	return &v1.RunnerListResp{
		Total: count,
		Data:  data,
	}, nil
}

func (w *WebServer) runnerListItemDO2VO(in *storage.RunnerModel, labels []string) *v1.RunnerListItem {
	return &v1.RunnerListItem{
		Id:        in.Id,
		Code:      in.Code,
		Name:      in.Name,
		Labels:    labels,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
