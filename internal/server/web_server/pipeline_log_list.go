package web_server

import (
	"context"
	"github.com/hl540/malou/internal/server/storage"
	v1 "github.com/hl540/malou/proto/v1"
)

func (w *WebServer) PipelineLogList(ctx context.Context, req *v1.PipelineLogListReq) (*v1.PipelineLogListResp, error) {
	logs, err := storage.GetPipelineLogs(ctx, req.PipelineId, req.Offset)
	if err != nil {
		return nil, err
	}
	var results []*v1.PipelineLog
	for _, log := range logs {
		results = append(results, &v1.PipelineLog{
			PipelineId: log.PipelineID,
			Type:       v1.PipelineLogType(v1.PipelineLogType_value[log.Type]),
			Step:       log.Step,
			Cmd:        log.Cmd,
			Message:    log.Message,
			Timestamp:  log.Timestamp,
			Duration:   log.Duration,
		})
	}
	return &v1.PipelineLogListResp{Log: results}, nil
}
