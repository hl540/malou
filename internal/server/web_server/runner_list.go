package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (w *WebServer) RunnerList(ctx context.Context, req *v1.RunnerListReq) (*v1.RunnerListResp, error) {
	list, count, err := storage.Runner.GetList(ctx, req.Code, req.Name, req.Labels, req.Page, req.Size)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	data := make([]*v1.RunnerListItem, 0, len(list))
	for _, v := range list {
		code, err := storage.Runner.GetLabelByRunnerID(ctx, v.Code)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		data = append(data, w.runnerLisDO2VO(v, w.runnerLabelDO2VO(code)))
	}
	return &v1.RunnerListResp{
		Total: count,
		Data:  data,
	}, nil
}

func (w *WebServer) runnerLisDO2VO(in *storage.RunnerModel, labels []string) *v1.RunnerListItem {
	return &v1.RunnerListItem{
		Code:      in.Code,
		Name:      in.Name,
		Labels:    labels,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
